package releaser

import (
	"github.com/commitsar-app/release-notary/internal/github"
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
}

// CreateReleaser initializes an instance of Releaser
func CreateReleaser(options Options) *Releaser {
	if options.Provider == "github" {
		githubService := github.CreateGithubService(options.Token, options.Owner, options.Repo)
		return &Releaser{service: githubService}
	}

	return &Releaser{}
}
