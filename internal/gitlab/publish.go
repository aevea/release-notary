package gitlab

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/commitsar-app/release-notary/internal/release"
)

// Publish publishes a Release https://developer.github.com/v3/repos/releases/#edit-a-release
func (g *Gitlab) Publish(release *release.Release) error {
	url := fmt.Sprintf("%v/projects/%v/repository/tags/%v/release", g.apiURL, g.projectID, g.tagName)

	jsonBody, err := json.Marshal(gitlabRelease{Message: release.Message})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := g.client.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Non 200 code received")
	}

	return nil
}
