package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/config"
	"sort"
)

// Contributions displays information about the repos, based on the number of PR/forks each has
func Contributions(conf *config.Config) {
	for _, res := range getReposSortedByContributions(conf) {
		fmt.Println(res)
	}
}

func getReposSortedByContributions(conf *config.Config) []string {
	repos := getReposWithPullRequests(conf)

	for i := range repos {
		repos[i].PrsPerForkCount = repos[i].ComputeContributionRatio()
	}

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].PrsPerForkCount > repos[j].PrsPerForkCount
	})

	var results []string

	for i := 0; i < conf.Top; i++ {
		results = append(results, repos[i].StringWithPullRequestsAndContributions())
	}

	return results
}
