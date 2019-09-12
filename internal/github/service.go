package github

import (
	"net/http"
)

// Github is the struct for mapping all functionality related to Github service
type Github struct {
	owner  string
	repo   string
	client *http.Client
}

type githubRelease struct {
	ID      int    `json:"id"`
	Tag     string `json:"tag_name"`
	Name    string `json:"name"`
	Message string `json:"body"`
}

// CreateGithubService initializes HTTP client and sets repo owner and repo name
func CreateGithubService(token, owner, repo string) *Github {

	return &Github{
		owner:  owner,
		repo:   repo,
		client: createClient(token),
	}
}
