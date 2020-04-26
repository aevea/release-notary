package gitlab

import (
	"fmt"

	"github.com/aevea/release-notary/internal/release"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// LatestRelease gets the current tag. https://docs.gitlab.com/ee/api/tags.html#get-a-single-repository-tag Gitlab does not have support for latest release in API.
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

	var glTag gitlabTag

	err = json.NewDecoder(response.Body).Decode(&glTag)

	if err != nil {
		return nil, err
	}

	message := ""

	// If a release already exists it's hidden in this field
	if glTag.Release.Message != "" {
		message = glTag.Release.Message
	}

	return &release.Release{ID: g.projectID, Tag: g.tagName, Name: glTag.Name, Message: message}, nil
}
