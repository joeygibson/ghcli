package commands

import (
	"github.com/joeygibson/ghcli/pkg/config"
	"math"
	"testing"
)

func TestSuccessfulFetchForContributions(t *testing.T) {
	setup(t)
	defer cleanup()

	conf := &config.Config{
		Org: "netflix",
		Top: 10,
	}

	repos := getReposSortedByContributions(conf)

	if len(repos) != conf.Top {
		t.Errorf("wrong result count; expected: %d, got %d", conf.Top, len(repos))
	}

	data := []struct {
		Index           int
		Name            string
		PrsPerForkCount float64
	}{
		{0, "astyanax", 0.021798},
		{1, "archaius", 0.014245},
		{2, "ribbon", 0.004202},
	}

	for _, d := range data {
		repo := repos[d.Index]

		if repo.Name != d.Name {
			t.Errorf("wrong repo at %d; expected %s, got: %s", d.Index, d.Name, repo.Name)
		}

		if !floatEquals(repo.PrsPerForkCount, d.PrsPerForkCount) {
			t.Errorf("wrong pr/fork count for %s; expected: %f, got: %f", repo.Name, d.PrsPerForkCount, repo.PrsPerForkCount)
		}
	}
}

var TOLERANCE = 0.001

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < TOLERANCE
}
