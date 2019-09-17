package gitlab

import (
	"fmt"

	"github.com/commitsar-app/release-notary/internal/release"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// LatestRelease gets the current tag. https://docs.gitlab.com/ee/api/releases/#get-a-release-by-a-tag-name Gitlab does not have support for latest release in API.
func (g *Gitlab) LatestRelease() (*release.Release, error) {
	url := fmt.Sprintf("%v/projects/%v/repository/tags/%v", g.apiURL, g.projectID, g.tagName)
	response, err := g.client.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%v returned %v code with error: %v", url, response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	var glRelease gitlabRelease

	err = json.NewDecoder(response.Body).Decode(&glRelease)

	if err != nil {
		return nil, err
	}

	return &release.Release{ID: g.projectID, Tag: g.tagName, Name: glRelease.Name, Message: glRelease.Message}, nil
}
