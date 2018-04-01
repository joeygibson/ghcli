package github

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type Repos []Repo

type Repo struct {
	Id               int    `json:"id"`
	Owner            Owner  `json:"owner"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Stargazers       int    `json:"stargazers_count"`
	Forks            int    `json:"forks_count"`
	PullRequestUrl   string `json:"pulls_url"`
	PullRequestCount int    `json:"-"`
}

type Owner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}

func (repo Repo) String() string {
	var buf strings.Builder

	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Name:\t%s\n", repo.Name)
	fmt.Fprintf(w, "Description:\t%s\n", repo.Description)
	fmt.Fprintf(w, "Stars:\t%d\n", repo.Stargazers)
	fmt.Fprintf(w, "Forks:\t%d\n", repo.Forks)

	if repo.PullRequestCount != 0 {
		fmt.Fprintf(w, "Pull requests:\t%d\n", repo.PullRequestCount)
	}

	w.Flush()

	return buf.String()
}
