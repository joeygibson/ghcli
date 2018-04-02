package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/client"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/joeygibson/ghcli/pkg/github"
	"github.com/sirupsen/logrus"
	"sort"
)

// Forks displays information about the repos, based on the number of forks each has
func Forks(conf *config.Config) {
	for _, repo := range getReposSortedByForks(conf) {
		fmt.Println(repo)
	}
}

func getReposSortedByForks(conf *config.Config) github.Repos {
	cl := client.New(conf)

	repos, err := cl.GetReposForOrg(conf.Org)
	if err != nil {
		logrus.Fatalf("getting repos for %s: %v", conf.Org, err)
	}

	logrus.Debugf("Repo count: %d\n", len(repos))

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].Forks > repos[j].Forks
	})

	if conf.Top >= len(repos) {
		return repos
	} else {
		return repos[0:conf.Top]
	}
}
