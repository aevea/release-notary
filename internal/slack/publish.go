package slack

import (
	"bytes"

	"github.com/aevea/quoad"
	jsoniter "github.com/json-iterator/go"
	"github.com/slack-go/slack"
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

	msg := slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: releaseNotes,
		},
	}

	return slack.PostWebhook(s.WebHookURL, &msg)
}
