package handler

import (
	"net/http"

	"github.com/RangHo/jumpstart/internal/artifact"
)

// HandleArtifact handles the request for a specific format of artifact.
func HandleArtifact(w http.ResponseWriter, r *http.Request, format string) {
	stream := r.URL.Query().Get("stream")
	architecture := r.URL.Query().Get("arch")

	// Find the artifact
	artifact, err := artifact.FindArtifact(stream, architecture, format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the artifact
	http.Redirect(w, r, artifact.ContentURL, http.StatusTemporaryRedirect)
}
