package version

import (
	"encoding/json"
)

// Kinds of Fedora CoreOS streams
const (
	Stable  = "stable"
	Testing = "testing"
	Next    = "next"
)

// Kinds of bare metal architectures
const (
	AMD64   = "x86_64"
	AArch64 = "aarch64"
	IBMZ    = "s390x"
	PowerPC = "ppc64le"
)

// Kinds of bare metal formats
const (
	ISO   = "iso"
	PXE   = "pxe"
	Raw   = "raw.xz"
	Raw4K = "4k.raw.xz"
)

type Artifact struct {
	Release      string
	Stream       string
	Architecture string
	Format       string
	ContentURL   string
	SignatureURL string
	Checksum     string
}

type File struct {
	Location string
	SHA256   string
}

func Validate(stream string, arch string) bool {
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

func Parse(input []byte) ([]Artifact, error) {
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
