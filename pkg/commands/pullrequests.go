package commands

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/client"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/joeygibson/ghcli/pkg/github"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

// PullRequests displays information about the repos, based on the number of PRs each has
func PullRequests(conf *config.Config) {
	for _, res := range getReposSortedByPullRequests(conf) {
		fmt.Println(res)
	}
}
func getReposWithPullRequests(conf *config.Config) github.Repos {
	cl := client.New(conf)

	repos, err := cl.GetReposForOrg(conf.Org)
	if err != nil {
		logrus.Fatalf("getting repos for %s: %v", conf.Org, err)
	}

	logrus.Debugf("Repo count: %d\n", len(repos))

	for i := range repos {
		url := strings.TrimSuffix(repos[i].PullRequestUrl, "{/number}")
		logrus.Debugf("PR URL: %s", url)

		prs, err := cl.GetPullRequestsForRepo(url)
		if err != nil {
			logrus.Errorf("getting pull requests for %s: %v", repos[i].Name, err)
			continue
		}

		repos[i].PullRequestCount = len(prs)
	}

	return repos
}

func getReposSortedByPullRequests(conf *config.Config) []string {
	repos := getReposWithPullRequests(conf)

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].PullRequestCount > repos[j].PullRequestCount
	})

	var results []string

	for i := 0; i < conf.Top; i++ {
		results = append(results, repos[i].StringWithPullRequests())
	}

	return results
}
