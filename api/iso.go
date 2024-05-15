package handler

import (
	"net/http"

	"github.com/RangHo/jumpstart"
)

func ISOHandler(w http.ResponseWriter, r *http.Request) {
	jumpstart.HandleArtifact(w, r, jumpstart.ISO)
}
