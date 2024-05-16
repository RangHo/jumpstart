package api

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/handler"
)

func InitramfsHandler(w http.ResponseWriter, r *http.Request) {
	handler.HandleArtifact(w, r, "pxe-initramfs")
}
