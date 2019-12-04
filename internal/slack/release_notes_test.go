package slack

import (
	"testing"

	"github.com/outillage/quoad"
	"github.com/stretchr/testify/assert"
)

func TestGenerateReleaseNotes(t *testing.T) {
	testData := map[string][]quoad.Commit{
		"features": []quoad.Commit{quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []quoad.Commit{quoad.Commit{Category: "chore", Scope: "", Heading: "testing"}, quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []quoad.Commit{quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"}, quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   []quoad.Commit{quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, quoad.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	expectedOutput := WebhookMessage(WebhookMessage{
		Blocks: []Block{
			Block{Type: "section", Section: content{Type: "mrkdwn", Text: "*Features*\r\nci test\r\n"}},
			Block{Type: "section", Section: content{Type: "mrkdwn", Text: "*Bug fixes*\r\nhuge bug\r\nbug fix\r\n"}},
			Block{Type: "section", Section: content{Type: "mrkdwn", Text: "*Chores and Improvements*\r\ntesting\r\nthis should end up in chores\r\n"}},
			Block{Type: "section", Section: content{Type: "mrkdwn", Text: "*Other*\r\nmerge master in something\r\nrandom\r\n"}}},
	})

	assert.Equal(t, expectedOutput, GenerateReleaseNotes(testData))
}
