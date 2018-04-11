package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/joeygibson/ghcli/pkg/github"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	conf      *config.Config
	githubUrl string
	client    *http.Client
}

const OrgRepoUrl = "https://api.github.com/orgs/%s/repos"

// New creates a structure for communicating with Github
func New(inConf *config.Config) Client {
	url := OrgRepoUrl

	specifiedUrl := os.Getenv("GHCLI_GITHUB_URL")
	if specifiedUrl != "" {
		url = specifiedUrl
	}

	return Client{
		conf:      inConf,
		githubUrl: url,
		client:    &http.Client{},
	}
}

func (c *Client) GetReposForOrg(org string) (github.Repos, error) {
	url := fmt.Sprintf(c.githubUrl, c.conf.Org)

	results, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var (
		allRepos github.Repos
		repos    github.Repos
	)

	for _, result := range results {
		err = json.Unmarshal(result, &repos)
		if err != nil {
			return nil, err
		}

		allRepos = append(allRepos, repos)
	}

	return allRepos, nil
}

func (c *Client) GetPullRequestsForRepo(url string) (github.PullRequests, error) {
	results, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var pullRequests github.PullRequests

	err = json.Unmarshal(result, &pullRequests)
	if err != nil {
		return nil, err
	}

	return pullRequests, nil
}

func (c *Client) get(getUrl string) ([][]byte, error) {
	var results [][]byte

	if len(strings.TrimSpace(getUrl)) == 0 {
		return nil, errors.New("invalid URL")
	}

	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return nil, err
	}

	if c.conf.Token != "" {
		logrus.Debug("Authenticating with OAuth token")
		authHeader := fmt.Sprintf("token %s", c.conf.Token)
		req.Header.Add("Authorization", authHeader)
	} else {
		logrus.Debug("Proceeding as anonymous")
	}

	for {
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("getting %s: %v", getUrl, err))
		}

		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			results = append(results, body)

			resp.Body.Close()

			nextUrl := getNextPageUrl(resp)

			if nextUrl == "" {
				return results, nil
			}

			req.URL, err = url.Parse(nextUrl)
			if err != nil {
				return results, err
			}

			continue
		}

		if resp.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized; either your token is invalid, or you used an incorrect user/password")
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.New("the org/repo you requested was not found")
		}

		if resp.StatusCode == http.StatusForbidden {
			return nil, errors.New("you are not permitted to access that org/repo")
		}
	}

	return nil, errors.New(fmt.Sprintf("the Github server was unable to complete your request: %v", err))
}

func getNextPageUrl(resp *http.Response) string {
	link := resp.Header.Get("Link")

	chunks := strings.Split(link, ",")

	for _, chunk := range chunks {
		chunks := strings.Split(chunk, ";")

		if chunks[1] == `rel="next"` {
			return strings.Trim(chunks[0], "<>")
		}
	}

	return ""
}
