package gitlab

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/aevea/release-notary/internal/release"
)

// Publish publishes a Release https://docs.gitlab.com/ee/api/releases/#create-a-release
func (g *Gitlab) Publish(release *release.Release) error {
	// By default, we are creating a new release
	method := "POST"
	url := fmt.Sprintf("%v/projects/%v/releases", g.apiURL, g.projectID)

	// In case release already exists we need to update instead of creating
	if release.Message != "" {
		url = fmt.Sprintf("%v/projects/%v/releases/%v", g.apiURL, g.projectID, g.tagName)
		method = "PUT"
	}

	jsonBody, err := json.Marshal(gitlabRelease{Message: release.ReleaseNotes, TagName: g.tagName})

	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := g.client.Do(req)

	if err != nil {
		return err
	}

	// 201 is used when a new release is attached to an existing tag
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return fmt.Errorf("%v %v returned %v code with error: %v", method, url, response.StatusCode, response.Status)
	}

	return nil
}
