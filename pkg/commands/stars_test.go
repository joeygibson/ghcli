package commands

import (
	"github.com/joeygibson/ghcli/pkg/config"
	"testing"
)

func TestSuccessfulFetchForStars(t *testing.T) {
	setup(t)
	defer cleanup()

	conf := &config.Config{
		Org: "netflix",
		Top: 10,
	}

	repos := getReposSortedByStars(conf)

	if len(repos) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(repos))
	}

	data := []struct {
		Index int
		Name  string
		Stars int
	}{
		{0, "Hystrix", 13054},
		{1, "SimianArmy", 6321},
		{6, "curator", 1816},
	}

	for _, d := range data {
		repo := repos[d.Index]

		if repo.Name != d.Name {
			t.Errorf("wrong repo at %d; expected %s, got: %s", d.Index, d.Name, repo.Name)
		}

		if repo.Stargazers != d.Stars {
			t.Errorf("wrong star count for %s; expected: %d, got: %d", repo.Name, d.Stars, repo.Stargazers)
		}
	}
}

func TestTopGreaterThanLength(t *testing.T) {
	setup(t)
	defer cleanup()

	conf := &config.Config{
		Org: "netflix",
		Top: 99999,
	}

	repos := getReposSortedByStars(conf)

	if len(repos) != 30 {
		t.Errorf("wrong result count; expected: %d, got %d", 30, len(repos))
	}
}
