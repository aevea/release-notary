package gitlab

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/commitsar-app/release-notary/internal/release"
)

// Publish publishes a Release https://developer.github.com/v3/repos/releases/#edit-a-release
func (g *Gitlab) Publish(release *release.Release) error {
	// By default we are creating a new release
	method := "POST"

	_, err := g.LatestRelease()

	if err != nil && err != errReleaseNotFound {
		return err
	}

	// In case release already exists we need to update instead of creating
	if err == errReleaseNotFound {
		method = "PUT"
	}

	url := fmt.Sprintf("%v/projects/%v/repository/tags/%v/release", g.apiURL, g.projectID, g.tagName)

	jsonBody, err := json.Marshal(gitlabRelease{Message: release.Message})

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

	if response.StatusCode != 200 {
		return fmt.Errorf("%v returned %v code with error: %v", url, response.StatusCode, response.Status)
	}

	return nil
}
