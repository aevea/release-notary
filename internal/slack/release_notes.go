package slack

import (
	"strings"

	"github.com/outillage/release-notary/internal/text"
)

// GenerateReleaseNotes creates a string from release notes that conforms with the Slack formatting. Expected format can be found in testdata.
func GenerateReleaseNotes(sections map[string][]text.Commit) string {
	builder := strings.Builder{}

	if len(sections["features"]) > 0 {
		builder.WriteString("*Features*\r\n")
		builder.WriteString(buildSection(sections["features"]))
		builder.WriteString("\r\n")
	}

	if len(sections["bugs"]) > 0 {
		builder.WriteString("*Bug fixes*\r\n")
		builder.WriteString(buildSection(sections["bugs"]))
		builder.WriteString("\r\n")
	}

	if len(sections["chores"]) > 0 {
		builder.WriteString("*Chores and Improvements*\r\n")
		builder.WriteString(buildSection(sections["chores"]))
		builder.WriteString("\r\n")
	}

	if len(sections["others"]) > 0 {
		builder.WriteString("*Other*\r\n")
		builder.WriteString(buildSection(sections["others"]))
		builder.WriteString("\r\n")
	}

	return builder.String()
}

func buildSection(commits []text.Commit) string {
	builder := strings.Builder{}

	for _, commit := range commits {
		builder.WriteString(commit.Heading)
		builder.WriteString("\r\n")
	}

	return builder.String()
}
