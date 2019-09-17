package gitlab

import (
	"errors"
	"net/http"
)

// Gitlab is a wrapper for the Gitlab service
type Gitlab struct {
	projectID int
	apiURL    string
	tagName   string
	client    *http.Client
}

type gitlabRelease struct {
	Name    string `json:"name"`
	Message string `json:"description"`
}

// CreateGitlabService creates an instance of the Gitlab service struct
func CreateGitlabService(projectID int, apiURL, tagName, token string) (*Gitlab, error) {
	if projectID == 0 {
		return nil, errors.New("missing projectID")
	}

	if apiURL == "" {
		return nil, errors.New("missing apiURL")
	}

	if tagName == "" {
		return nil, errors.New("missing tagName")
	}

	if token == "" {
		return nil, errors.New("missing token")
	}

	return &Gitlab{projectID: projectID, apiURL: apiURL, tagName: tagName, client: createClient(token)}, nil
}
