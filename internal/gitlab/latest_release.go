package gitlab

import (
	"errors"
	"fmt"

	"github.com/commitsar-app/release-notary/internal/release"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// LatestRelease gets the release on current tag. https://docs.gitlab.com/ee/api/releases/#get-a-release-by-a-tag-name Gitlab does not have support for latest release in API.
func (g *Gitlab) LatestRelease() (*release.Release, error) {
	url := fmt.Sprintf("%v/projects/%v/releases/%v", g.apiURL, g.projectID, g.tagName)
	response, err := g.client.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("api returned non 200 response")
	}

	defer response.Body.Close()

	var glRelease gitlabRelease

	err = json.NewDecoder(response.Body).Decode(&glRelease)

	if err != nil {
		return nil, err
	}

	return &release.Release{ID: g.projectID, Tag: g.tagName, Name: glRelease.Name, Message: glRelease.Message}, nil
}
