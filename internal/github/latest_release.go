package github

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/commitsar-app/release-notary/internal/release"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// LatestRelease calls https://developer.github.com/v3/repos/releases/#get-the-latest-release and returns a Release struct
func (g *Github) LatestRelease() (*release.Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/releases/latest", g.owner, g.repo)
	response, err := g.client.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("hmm")
	}

	defer response.Body.Close()

	var ghRelease githubRelease

	err = json.NewDecoder(response.Body).Decode(&ghRelease)

	if err != nil {
		return nil, err
	}


	return &release.Release{ID: ghRelease.ID, Tag: ghRelease.Tag, Name: ghRelease.Name, Message: ghRelease.Message}, nil
}
