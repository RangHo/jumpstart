package api

import (
	"net/http"
)

const repoURL = "https://github.com/RangHo/jumpstart"

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, repoURL, http.StatusTemporaryRedirect)
}
