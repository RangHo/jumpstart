package pxehandler

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/artifact"
	"github.com/RangHo/jumpstart/pkg/version"
)

func KernelHandler(w http.ResponseWriter, r *http.Request) {
	artifact.Handle(w, r, version.PXE+"-kernel")
}
