package github

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/commitsar-app/release-notary/internal/release"
)

// Publish publishes a Release https://developer.github.com/v3/repos/releases/#edit-a-release
func (g *Github) Publish(release *release.Release) error {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/releases/%v", g.owner, g.repo, release.ID)

	jsonBody, err := json.Marshal(githubRelease{ID: release.ID, Tag: release.Tag, Name: release.Name, Message: release.Message})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))

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
