package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joeygibson/ghcli/pkg/github"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	server *httptest.Server
)

func setup(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.RequestURI, "/pulls") {
			// return dummied-up PRs for the given repo
			chunks := strings.Split(r.RequestURI, "/")
			repo := chunks[3]

			var prs github.PullRequests

			if repo == "archaius" {
				prs = github.PullRequests{
					{Id: 1, State: "open"},
					{Id: 2, State: "open"},
					{Id: 3, State: "open"},
					{Id: 4, State: "open"},
					{Id: 5, State: "open"},
				}
			} else if repo == "astyanax" {
				prs = github.PullRequests{
					{Id: 1, State: "open"},
					{Id: 2, State: "open"},
					{Id: 3, State: "open"},
					{Id: 4, State: "open"},
					{Id: 5, State: "open"},
					{Id: 6, State: "open"},
					{Id: 7, State: "open"},
					{Id: 8, State: "open"},
				}
			} else if repo == "ribbon" {
				prs = github.PullRequests{
					{Id: 1, State: "open"},
					{Id: 2, State: "open"},
				}
			}

			result, _ := json.Marshal(prs)

			w.WriteHeader(http.StatusOK)
			io.Copy(w, bytes.NewReader(result))
		} else {
			// Return the whole list of repos for the organization
			fileName := "testdata/repos.json"

			fileContents, err := ioutil.ReadFile(fileName)
			if err != nil {
				t.Errorf("loading %s: %v", fileName, err)
			}

			contents := string(fileContents)

			newUrl := fmt.Sprintf("http://%s", r.Host)
			contents = strings.Replace(contents, "https://api.github.com", newUrl, -1)

			w.WriteHeader(http.StatusOK)
			io.Copy(w, bytes.NewReader([]byte(contents)))
		}
	}

	server = httptest.NewServer(http.HandlerFunc(handler))

	url := server.URL + "/orgs/%s/repos"
	os.Setenv("GHCLI_GITHUB_URL", url)
}

func cleanup() {
	server.Close()
}
