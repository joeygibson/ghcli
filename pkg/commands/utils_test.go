package commands

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server *httptest.Server
)

func setup(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fileName := "testdata/repos.json"

		fileContents, err := ioutil.ReadFile(fileName)
		if err != nil {
			t.Errorf("loading %s: %v", fileName, err)
		}

		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(fileContents))
	}

	server = httptest.NewServer(http.HandlerFunc(handler))

	url := server.URL + "/orgs/%s/repos"
	os.Setenv("GHCLI_GITHUB_URL", url)
}

func cleanup() {
	server.Close()
}
