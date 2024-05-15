package passwdhandler

import (
	"net/http"

	"github.com/RangHo/jumpstart"
)

func PasswdFromGitHubHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	ignition, err := jumpstart.MakePasswdFromGitHub(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ignition)
}
