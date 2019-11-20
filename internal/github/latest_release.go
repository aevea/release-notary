package github

import (
	"fmt"

	"github.com/commitsar-app/release-notary/internal/release"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// LatestRelease calls https://developer.github.com/v3/repos/releases/#get-the-latest-release and returns a Release struct
func (g *Github) LatestRelease() (*release.Release, error) {
	url := fmt.Sprintf("%v/%v/%v/releases/latest", g.APIURL, g.owner, g.repo)
	response, err := g.client.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%v returned %v code with error: %v", url, response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	var ghRelease githubRelease

	err = json.NewDecoder(response.Body).Decode(&ghRelease)

	if err != nil {
		return nil, err
	}

	return &release.Release{ID: ghRelease.ID, Tag: ghRelease.Tag, Name: ghRelease.Name, Message: ghRelease.Message}, nil
}
