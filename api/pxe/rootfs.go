package api

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/handler"
)

func RootfsHandler(w http.ResponseWriter, r *http.Request) {
	handler.HandleArtifact(w, r, "pxe-rootfs")
}
