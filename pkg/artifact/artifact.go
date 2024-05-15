package artifact

import (
	"fmt"
	"io"
	"net/http"

	"github.com/RangHo/jumpstart/pkg/version"
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
	parsed, err := version.Parse(body)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func Find(stream string, architecture string, format string) (string, error) {
	// Validate the stream and architecture
	if !version.Validate(stream, architecture) {
		return "", fmt.Errorf("Invalid stream or architecture")
	}

	// Fetch the artifacts
	artifacts, err := getArtifacts(stream)
	if err != nil {
		return "", err
	}

	// Find the artifact
	for _, artifact := range artifacts {
		if artifact.Architecture == architecture && artifact.Format == format {
			return artifact.ContentURL, nil
		}
	}

	return "", fmt.Errorf("Artifact not found")
}

func Handle(w http.ResponseWriter, r *http.Request, format string) {
	stream := r.URL.Query().Get("stream")
	architecture := r.URL.Query().Get("architecture")

	// Find the artifact
	url, err := Find(stream, architecture, format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Redirect to the artifact
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
