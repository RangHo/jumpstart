package artifact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// baseJSON is the base URL for the Fedora CoreOS stream metadata.
const baseJSON = "https://builds.coreos.fedoraproject.org/streams/%s.json"

// Fedora CoreOS provides these streams for updates.
const (
	Stable  = "stable"
	Testing = "testing"
	Next    = "next"
)

// Fedora CoreOS provides these architectures for bare-metal installation.
const (
	AMD64   = "x86_64"
	AArch64 = "aarch64"
	IBMZ    = "s390x"
	PowerPC = "ppc64le"
)

// Fedora CoreOS provides these formats for bare-metal installation.
const (
	ISO   = "iso"
	PXE   = "pxe"
	Raw   = "raw.xz"
	Raw4K = "4k.raw.xz"
)

// Artifact represents a Fedora CoreOS artifact that is accessible through Jumpstart.
type Artifact struct {
	Release      string // release version
	Stream       string // release stream
	Architecture string // target architecture
	Format       string // image format
	ContentURL   string // URL to the artifact content
	SignatureURL string // URL to the artifact signature
	Checksum     string // SHA256 checksum of the artifact content
}

// rawStream represents the whole Fedora CoreOS stream JSON document.
type rawStream struct {
	Stream   string `json:"stream"`
	Metadata struct {
		LastModified time.Time `json:"last-modified"`
		Generator    string    `json:"generator"`
	}
	Architectures map[string]architectureValue `json:"architectures"`
}

// architectureValue represents the architecture-specific part of the Fedora CoreOS stream JSON document.
//
// While there are plethora of artifact types available, but we're interested in the bare-metal artifacts only.
type architectureValue struct {
	Artifacts struct {
		Metal artifactValue `json:"metal"`
	} `json:"artifacts"`
}

// artifactValue represents the artifact-specific part of the Fedora CoreOS stream JSON document.
type artifactValue struct {
	Release string                 `json:"release"`
	Formats map[string]formatValue `json:"formats"`
}

// formatValue represents the format-specific part of the Fedora CoreOS stream JSON document.
type formatValue map[string]remoteFile

// remoteFile represents the remote file information.
type remoteFile struct {
	Location           string  `json:"location"`
	Signature          string  `json:"signature"`
	SHA256             string  `json:"sha256"`
	UncompressedSHA256 *string `json:"uncompressed-sha256"`
}

// parse parses the Fedora CoreOS stream JSON document into a list of artifacts.
func parse(input []byte) ([]Artifact, error) {
	raw := rawStream{}
	err := json.Unmarshal(input, &raw)
	if err != nil {
		return nil, err
	}

	knownInfo := Artifact{
		Stream: raw.Stream,
	}

	result := []Artifact{}

	for arch, archValue := range raw.Architectures {
		knownInfo.Architecture = arch
		parsed, err := parseArchitectures(archValue, knownInfo)
		if err != nil {
			return nil, err
		}

		result = append(result, parsed...)
	}

	return result, nil
}

// parseArchitectures parses the architecture-specific part of the Fedora CoreOS stream JSON document.
func parseArchitectures(input architectureValue, info Artifact) ([]Artifact, error) {
	info.Release = input.Artifacts.Metal.Release

	result := []Artifact{}

	for format, kind := range input.Artifacts.Metal.Formats {
		info.Format = format
		parsed, err := parseFormats(kind, info)
		if err != nil {
			return nil, err
		}

		result = append(result, parsed...)
	}

	return result, nil
}

// parseFormats parses the format-specific part of the Fedora CoreOS stream JSON document.
func parseFormats(input formatValue, info Artifact) ([]Artifact, error) {
	result := []Artifact{}

	switch info.Format {
	case Raw, Raw4K, ISO:
		result = append(result, Artifact{
			Release:      info.Release,
			Stream:       info.Stream,
			Architecture: info.Architecture,
			Format:       info.Format,
			ContentURL:   input["disk"].Location,
			SignatureURL: input["disk"].Signature,
			Checksum:     input["disk"].SHA256,
		})
	case PXE:
		parsed, err := parsePXE(input, info)
		if err != nil {
			return nil, err
		}

		result = append(result, parsed...)
	}

	return result, nil
}

// parsePXE parses the PXE-specific part of the Fedora CoreOS stream JSON document.
func parsePXE(input formatValue, info Artifact) ([]Artifact, error) {
	result := []Artifact{}

	for kind, file := range input {
		result = append(result, Artifact{
			Release:      info.Release,
			Stream:       info.Stream,
			Architecture: info.Architecture,
			Format:       info.Format + "-" + kind,
			ContentURL:   file.Location,
			SignatureURL: file.Signature,
			Checksum:     file.SHA256,
		})
	}

	return result, nil
}

// getArtifacts fetches the Fedora CoreOS stream metadata and parses it into a list of artifacts.
func getArtifacts(stream string) ([]Artifact, error) {
	// Fetch the stream metadata
	resp, err := http.Get(fmt.Sprintf(baseJSON, stream))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the content
	parsed, err := parse(body)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

// validate checks if the stream and architecture are valid.
func validate(stream string, arch string) bool {
	switch stream {
	case Stable, Testing, Next:
		// Do nothing
	default:
		return false
	}

	switch arch {
	case AMD64, AArch64, IBMZ, PowerPC:
		// Do nothing
	default:
		return false
	}

	return true
}

// FindArtifact finds the artifact for the given stream, architecture, and format.
func FindArtifact(stream string, architecture string, format string) (Artifact, error) {
	// Validate the stream and architecture
	if !validate(stream, architecture) {
		return Artifact{}, fmt.Errorf("Invalid stream or architecture")
	}

	// Fetch the artifacts
	artifacts, err := getArtifacts(stream)
	if err != nil {
		return Artifact{}, err
	}

	// Find the artifact
	for _, artifact := range artifacts {
		if artifact.Architecture == architecture && artifact.Format == format {
			return artifact, nil
		}
	}

	return Artifact{}, fmt.Errorf("Artifact not found")
}
