package releaser

import (
	"fmt"

	"github.com/commitsar-app/release-notary/internal/github"
	"github.com/commitsar-app/release-notary/internal/gitlab"
	"github.com/commitsar-app/release-notary/internal/release"
)

// Service describes the basic usage needed from a service such as Github
type Service interface {
	LatestRelease() (*release.Release, error)
	Publish(*release.Release) error
}

// Releaser holds all functionality related to releasing.
type Releaser struct {
	service Service
}

// Options are used to initialize releaser
type Options struct {
	Provider Provider
	DryRun   bool
	Token    string
	Owner    string
	Repo     string
	// used for setting gitlab api url
	APIURL string
	// used by gitlab
	TagName string
	// used by Gitlab
	ProjectID int
}

// CreateReleaser initializes an instance of Releaser
func CreateReleaser(options Options) (*Releaser, error) {
	if options.Provider == "github" {
		githubService := github.CreateGithubService(options.Token, options.Owner, options.Repo)
		return &Releaser{service: githubService}, nil
	}

	if options.Provider == "gitlab" {
		gitlabService, err := gitlab.CreateGitlabService(options.ProjectID, options.APIURL, options.TagName, options.Token)
		return &Releaser{service: gitlabService}, err
	}

	return nil, fmt.Errorf("provider %v not found", options.Provider)
}
