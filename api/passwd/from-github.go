package api

import (
	"net/http"

	"github.com/RangHo/jumpstart/pkg/handler"
)

func PasswdFromGitHubHandler(w http.ResponseWriter, r *http.Request) {
	handler.HandlePasswdFromGitHub(w, r)
}
