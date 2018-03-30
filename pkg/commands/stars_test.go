package commands

import (
	"bytes"
	"github.com/joeygibson/ghcli/pkg/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuccessfulFetch(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fileName := "testdata/repos.json"

		fileContents, err := ioutil.ReadFile(fileName)
		if err != nil {
			t.Errorf("loading %s: %v", fileName, err)
		}

		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(fileContents))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	conf := &config.Config{
		Org: "netflix",
		Top: 10,
	}

	results := getReposByStars(conf)

	if len(results) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(results))
	}

	res := results[0]

	if !strings.Contains(res, "Name: Hystrix") {
		t.Errorf("wrong repo at top; expected Hystrix, got: %s", res)
	}

	if !strings.Contains(res, "Stars: 13054") {
		t.Errorf("wrong star count for Hystrix; expected: %d, got: %s", 13054, res)
	}
}
