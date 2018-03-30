package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/client"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/sirupsen/logrus"
	"sort"
)

// Stars displays information about the highest-starred repos for the given organization.
func Stars(conf *config.Config) {
	for _, res := range getReposByStars(conf) {
		fmt.Println(res)
	}
}

func getReposByStars(conf *config.Config) []string {
	cl := client.New(conf)

	repos, err := cl.GetReposForOrg(conf.Org)
	if err != nil {
		logrus.Fatalf("getting repos for %s: %v", conf.Org, err)
	}

	logrus.Debugf("Repo count: %d\n", len(repos))

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].Stargazers > repos[j].Stargazers
	})

	var results []string

	for i := 0; i < conf.Top; i++ {
		results = append(results, repos[i].String())
	}

	return results
}
