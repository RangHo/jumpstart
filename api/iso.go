package api

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/handler"
)

func ISOHandler(w http.ResponseWriter, r *http.Request) {
	handler.HandleArtifact(w, r, "iso")
}
