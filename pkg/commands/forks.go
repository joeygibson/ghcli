package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/client"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/sirupsen/logrus"
	"sort"
)

// Forks displays information about the repos, based on the number of forks each has
func Forks(conf *config.Config) {
	for _, res := range getReposByForks(conf) {
		fmt.Println(res)
	}
}

func getReposByForks(conf *config.Config) []string {
	cl := client.New(conf)

	repos, err := cl.GetReposForOrg(conf.Org)
	if err != nil {
		logrus.Fatalf("getting repos for %s: %v", conf.Org, err)
	}

	logrus.Debugf("Repo count: %d\n", len(repos))

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].Forks > repos[j].Forks
	})

	var results []string

	for i := 0; i < conf.Top; i++ {
		results = append(results, repos[i].String())
	}

	return results
}
