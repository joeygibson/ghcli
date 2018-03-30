package github

import "fmt"

type Repos []Repo

type Repo struct {
	Id          int    `json:"id"`
	Owner       Owner  `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stargazers  int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
}

type Owner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}

func (repo Repo) String() string {
	return fmt.Sprintf("Name: %s\nDescription: %s\nStars: %d\n", repo.Name,
		repo.Description, repo.Stargazers)
}
