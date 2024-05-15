package api

import (
	"net/http"

	"github.com/RangHo/jumpstart"
)

func RawHandler(w http.ResponseWriter, r *http.Request) {
	jumpstart.HandleArtifact(w, r, jumpstart.Raw)
}
