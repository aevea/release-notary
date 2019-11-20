package slack

import (
	"bytes"
	"net/http"
	"time"

	"github.com/commitsar-app/release-notary/internal/text"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type request struct {
	Text string `json:"text"`
}

// Publish pushes the release notes to Slack via provided Webhook. https://api.slack.com/reference/messaging/payload
func (s *Slack) Publish(commits map[string][]text.Commit) error {
	releaseNotes := GenerateReleaseNotes(commits)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	jsonBody, err := json.Marshal(request{Text: releaseNotes})

	if err != nil {
		return err
	}

	_, err = client.Post(s.WebHookURL, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	return nil
}
