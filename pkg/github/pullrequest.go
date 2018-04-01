package github

type PullRequests []PullRequest

type PullRequest struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	State string `json:"state"`
}
