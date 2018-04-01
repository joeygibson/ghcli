package commands

import (
	"github.com/joeygibson/ghcli/pkg/config"
	"testing"
)

func TestSuccessfulFetchForForks(t *testing.T) {
	setup(t)
	defer cleanup()

	conf := &config.Config{
		Org: "netflix",
		Top: 10,
	}

	repos := getReposByForks(conf)

	if len(repos) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(repos))
	}

	data := []struct {
		Index int
		Name  string
		Forks int
	}{
		{0, "Hystrix", 2611},
		{1, "eureka", 1190},
		{6, "asgard", 436},
	}

	for _, d := range data {
		repo := repos[d.Index]

		if repo.Name != d.Name {
			t.Errorf("wrong repo at top; expected %s, got: %s", d.Name, repo.Name)
		}

		if repo.Forks != d.Forks {
			t.Errorf("wrong fork count for top; expected: %d, got: %d", d.Forks, repo.Forks)
		}
	}
}
