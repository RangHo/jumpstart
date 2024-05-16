package handler

import (
	"net/http"

	"github.com/RangHo/jumpstart/internal/passwd"
)

// HandlePasswdFromGitHub handles the request for a password file generated from a GitHub username.
func HandlePasswdFromGitHub(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	ignition, err := passwd.MakePasswdFromGitHub(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ignition)
}
