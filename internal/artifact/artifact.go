package artifact

import (
	"fmt"
	"io"
	"net/http"

	"github.com/RangHo/jumpstart/internal/version"
)

// Base URL for fetching the latest Fedora CoreOS stream metadata
const baseJSON = "https://builds.coreos.fedoraproject.org/streams/%s.json"

// Parse the metadata from the stream
func getArtifacts(stream string) ([]version.Artifact, error) {
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
	return version.Parse(body), nil
}

func getPXE(stream string, architecture string, imageType string) (string, error) {
	// Validate the stream-arch pair
	if !version.Validate(stream, architecture) {
		return "", fmt.Errorf("Invalid stream-arch pair: %s-%s", stream, architecture)
	}

	// Fetch the artifacts
	artifacts, err := getArtifacts(stream)
	if err != nil {
		return "", err
	}

	// Find the artifact
	for _, artifact := range artifacts {
		if artifact.Architecture == architecture && artifact.Type == imageType {
			return artifact.Content.Location, nil
		}
	}

	return "", fmt.Errorf("Requested file not found for %s-%s", stream, architecture)
}

func GetKernel(stream string, architecture string) (string, error) {
	return getPXE(stream, architecture, version.PXE + "-kernel")
}

func GetInitramfs(stream string, architecture string) (string, error) {
	return getPXE(stream, architecture, version.PXE + "-initramfs")
}

func GetRootfs(stream string, architecture string) (string, error) {
	return getPXE(stream, architecture, version.PXE + "-rootfs")
}
