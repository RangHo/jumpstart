package handler

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/artifact"
	"github.com/RangHo/jumpstart/pkg/version"
)

func ISOHandler(w http.ResponseWriter, r *http.Request) {
	artifact.Handle(w, r, version.ISO)
}
