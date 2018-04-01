package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/joeygibson/ghcli/pkg/github"
	"sort"
)

// Contributions displays information about the repos, based on the number of PR/forks each has
func Contributions(conf *config.Config) {
	for _, res := range getReposSortedByContributions(conf) {
		fmt.Println(res.StringWithPullRequestsAndContributions())
	}
}

func getReposSortedByContributions(conf *config.Config) github.Repos {
	repos := getReposWithPullRequests(conf)

	for i := range repos {
		repos[i].PrsPerForkCount = repos[i].ComputeContributionRatio()
	}

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].PrsPerForkCount > repos[j].PrsPerForkCount
	})

	if conf.Top >= len(repos) {
		return repos
	} else {
		return repos[0:conf.Top]
	}
}
