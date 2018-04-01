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
	"strings"
)

type Client struct {
	conf   *config.Config
	client *http.Client
}

const OrgRepoUrl = "https://api.github.com/orgs/%s/repos"

// New creates a strucure for communicating with Github
func New(inConf *config.Config) Client {
	return Client{
		conf:   inConf,
		client: &http.Client{},
	}
}

func (c *Client) GetReposForOrg(org string) (github.Repos, error) {
	url := fmt.Sprintf(OrgRepoUrl, c.conf.Org)

	result, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var repos github.Repos

	err = json.Unmarshal(result, &repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (c *Client) GetPullRequestsForRepo(url string) (github.PullRequests, error) {
	result, err := c.get(url)
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

func (c *Client) get(url string) ([]byte, error) {
	if len(strings.TrimSpace(url)) == 0 {
		return nil, errors.New("invalid URL")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.conf.Token != "" {
		logrus.Debug("Authenticating with OAuth token")
		authHeader := fmt.Sprintf("token %s", c.conf.Token)
		req.Header.Add("Authorization", authHeader)
	} else if c.conf.User != "" && c.conf.Password != "" {
		logrus.Debug("Authenticating with basic auth")
		req.SetBasicAuth(c.conf.User, c.conf.Password)
	} else {
		logrus.Debug("Proceeding as anonymous")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting %s: %v", url, err))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
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

	return nil, errors.New(fmt.Sprintf("the Github server was unable to complete your request: %v", err))
}
