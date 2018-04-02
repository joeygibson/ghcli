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
	repos := getReposSortedByPullRequests(conf)

	for _, repo := range repos {
		fmt.Println(repo.StringWithPullRequests())
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

		prs, err := cl.GetPullRequestsForRepo(url)
		if err != nil {
			logrus.Errorf("getting pull requests for %s: %v", repos[i].Name, err)
			continue
		}

		repos[i].PullRequestCount = len(prs)
	}

	return repos
}

func getReposSortedByPullRequests(conf *config.Config) github.Repos {
	repos := getReposWithPullRequests(conf)

	sort.Slice(repos[:], func(i, j int) bool {
		return repos[i].PullRequestCount > repos[j].PullRequestCount
	})

	if conf.Top >= len(repos) {
		return repos
	} else {
		return repos[0:conf.Top]
	}
}
