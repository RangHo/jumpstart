package pxehandler

import (
	"net/http"

	"github.com/RangHo/jumpstart"
)

func InitramfsHandler(w http.ResponseWriter, r *http.Request) {
	jumpstart.HandleArtifact(w, r, jumpstart.PXE+"-initramfs")
}
