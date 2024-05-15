package pxehandler

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/artifact"
)

func RootfsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the stream and architecture from the parameters
	stream := r.URL.Query().Get("stream")
	architecture := r.URL.Query().Get("arch")

	url, err := artifact.GetRootfs(stream, architecture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
