package pxehandler

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/artifact"
	"github.com/RangHo/jumpstart/pkg/version"
)

func RootfsHandler(w http.ResponseWriter, r *http.Request) {
	stream := r.URL.Query().Get("stream")
	architecture := r.URL.Query().Get("arch")

	url, err := artifact.FindArtifact(stream, architecture, version.PXE + "-rootfs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
