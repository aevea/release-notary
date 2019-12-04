package slack

import (
	"bytes"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/outillage/quoad"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Publish pushes the release notes to Slack via provided Webhook. https://api.slack.com/reference/messaging/payload
func (s *Slack) Publish(commits map[string][]quoad.Commit) error {
	releaseNotes := GenerateReleaseNotes(commits)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	jsonBody, err := json.Marshal(releaseNotes)

	if err != nil {
		return err
	}

	_, err = client.Post(s.WebHookURL, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	return nil
}
