package indexhandler

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	for _, path := range paths {
		fmt.Fprintf(w, "Path: %s\n", path)
	}
}
