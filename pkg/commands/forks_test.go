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

func TestSuccessfulFetchForForks(t *testing.T) {
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

	results := getReposByForks(conf)

	if len(results) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(results))
	}

	res := results[0]
	forksLine := strings.Split(res, "\n")[3]
	forksChunks := strings.Split(forksLine, ":")
	forks := strings.TrimSpace()
	if !strings.Contains(res, "Hystrix") {
		t.Errorf("wrong repo at top; expected Hystrix, got: %s", res)
	}

	if !strings.Contains(res, "2611") {
		t.Errorf("wrong fork count for top; expected: %d, got: %s", 2611, res)
	}

	res = results[1]

	if !strings.Contains(res, "eureka") {
		t.Errorf("wrong repo at top; expected eureka, got: %s", res)
	}

	if !strings.Contains(res, "1190") {
		t.Errorf("wrong fork count for top; expected: %d, got: %s", 1190, res)
	}
}
