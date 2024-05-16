package pxe

import (
	"net/http"

	"github.com/RangHo/jumpstart"
)

func RootfsHandler(w http.ResponseWriter, r *http.Request) {
	jumpstart.HandleArtifact(w, r, jumpstart.PXE+"-rootfs")
}
