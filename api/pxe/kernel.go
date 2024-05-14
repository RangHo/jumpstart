package pxekernelhandler

import (
	"net/http"

	"github.com/RangHo/jumpstart/internal/artifact"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the stream and architecture from the parameters
	stream := r.URL.Query().Get("stream")
	architecture := r.URL.Query().Get("arch")

	url, err := artifact.GetKernel(stream, architecture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
