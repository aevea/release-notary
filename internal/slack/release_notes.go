package slack

import (
	"strings"

	"github.com/outillage/quoad"
)

// GenerateReleaseNotes creates a string from release notes that conforms with the Slack formatting. Expected format can be found in testdata.
func GenerateReleaseNotes(sections map[string][]quoad.Commit) WebhookMessage {
	var blocks []Block

	if len(sections["features"]) > 0 {
		section := Block{Type: "section", Section: buildSection("Features", sections["features"])}
		blocks = append(blocks, section)
	}

	if len(sections["bugs"]) > 0 {
		section := Block{Type: "section", Section: buildSection("Bug fixes", sections["bugs"])}
		blocks = append(blocks, section)
	}

	if len(sections["chores"]) > 0 {
		section := Block{Type: "section", Section: buildSection("Chores and Improvements", sections["chores"])}
		blocks = append(blocks, section)
	}

	if len(sections["others"]) > 0 {
		section := Block{Type: "section", Section: buildSection("Other", sections["others"])}
		blocks = append(blocks, section)
	}

	return WebhookMessage{Blocks: blocks}
}

func buildSection(heading string, commits []quoad.Commit) content {
	builder := strings.Builder{}
	builder.WriteString("*" + heading + "*\r\n")

	for _, commit := range commits {
		builder.WriteString(commit.Heading)
		builder.WriteString("\r\n")
	}

	section := content{Type: "mrkdwn", Text: builder.String()}

	return section
}
