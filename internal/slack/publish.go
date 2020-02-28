package slack

import (
	"bytes"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/outillage/quoad"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)

	return buffer.Bytes(), err
}

// Publish pushes the release notes to Slack via provided Webhook. https://api.slack.com/reference/messaging/payload
func (s *Slack) Publish(commits map[string][]quoad.Commit, remote GitRemoter) error {
	releaseNotes := GenerateReleaseNotes(commits, remote)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	jsonBody, err := jsonMarshal(releaseNotes)

	if err != nil {
		return err
	}

	_, err = client.Post(s.WebHookURL, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return err
	}

	return nil
}
