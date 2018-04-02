package commands

import (
	"github.com/joeygibson/ghcli/pkg/config"
	"testing"
)

func TestSuccessfulFetchForPullRequests(t *testing.T) {
	setup(t)
	defer cleanup()

	conf := &config.Config{
		Org: "netflix",
		Top: 10,
	}

	repos := getReposSortedByPullRequests(conf)

	if len(repos) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(repos))
	}

	data := []struct {
		Index            int
		Name             string
		PullRequestCount int
	}{
		{0, "astyanax", 8},
		{1, "archaius", 5},
		{2, "ribbon", 2},
	}

	for _, d := range data {
		repo := repos[d.Index]

		if repo.Name != d.Name {
			t.Errorf("wrong repo at %d; expected %s, got: %s", d.Index, d.Name, repo.Name)
		}

		if repo.PullRequestCount != d.PullRequestCount {
			t.Errorf("wrong pull request count for %s; expected: %d, got: %d", repo.Name, d.PullRequestCount, repo.PullRequestCount)
		}
	}
}
