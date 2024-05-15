package jumpstart

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

// passwdFromGithubBuTmpl is the Butane template for the passwd-from-github.bu config.
//
//go:embed configs/passwd-from-github.bu.tmpl
var passwdFromGithubBuTmpl []byte

// passwdFromGitHubData is the data structure for the passwd-from-github.bu config.
type passwdFromGitHubData struct {
	Name string
	Keys []string
}

// MakePasswdFromGitHub creates an Ignition config for the passwd-from-github.bu template.
func MakePasswdFromGitHub(name string) ([]byte, error) {
	res, err := http.Get("https://github.com/" + name + ".keys")
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("failed to fetch keys: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	keys := strings.Split(string(body), "\n")
	keys = keys[:len(keys)-1]

	tmpl, err := template.New("passwd-from-github.bu").Parse(string(passwdFromGithubBuTmpl))
	if err != nil {
		return []byte{}, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, passwdFromGitHubData{
		Name: "core",
		Keys: keys,
	})
	if err != nil {
		return []byte{}, err
	}

	return Ignite(buf.Bytes())
}
