package api

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/handler"
)

func RawHandler(w http.ResponseWriter, r *http.Request) {
	handler.HandleArtifact(w, r, "raw.xz")
}
