package slack

import (
	"fmt"
	"strings"

	"github.com/aevea/quoad"
	"github.com/aevea/release-notary/internal"
	"github.com/slack-go/slack"
)

func pluralize(base string, count int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d %s", count, base))

	if count != 1 {
		builder.WriteString("s")
	}

	return builder.String()
}

func countReferences(commits []quoad.Commit) int {
	count := 0

	for _, commit := range commits {
		count += len(commit.Issues)
	}

	return count
}

// GenerateReleaseNotes creates a string from release notes that conforms with the Slack formatting. Expected format can be found in testdata.
func GenerateReleaseNotes(changelog map[string][]quoad.Commit, remote GitRemoter) []slack.Block {
	blocks := []slack.Block{buildReleaseTitle(remote)}

	sections := internal.SortedSections

	currentSection := 0
	for _, section := range sections {
		commits := changelog[section.ID]
		if len(commits) > 0 {

			sectionTitle := slack.SectionBlock{
				Type: "section",
				Text: &slack.TextBlockObject{
					Type: "mrkdwn",
					Text: fmt.Sprintf(":%s: *%s*", section.Icon, section.Title),
				},
			}

			sectionContext := slack.NewContextBlock(
				"",
				&slack.TextBlockObject{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"%s referencing %s",
						pluralize("commit", len(commits)),
						pluralize("issue", countReferences(commits)),
					),
				},
			)

			blocks = append(
				blocks,
				sectionTitle,
				sectionContext,
			)

			blocks = append(blocks, buildCommitList(commits, remote.GetRemoteURL())...)

			// Check if there is another section following this one in order to display a divider
			if currentSection+1 < len(sections) {
				blocks = append(blocks, slack.NewDividerBlock())
			}

			currentSection++
		}
	}

	return blocks
}

func buildReleaseTitle(remote GitRemoter) slack.Block {
	return slack.HeaderBlock{
		Type: "header",
		Text: &slack.TextBlockObject{
			Type:  "plain_text",
			Text:  fmt.Sprintf(":tada: New release for %s", remote.Project()),
			Emoji: true,
		},
	}
}

func addScope(scope string) string {
	if scope == "" {
		return ""
	}
	return fmt.Sprintf("*(%s)* ", scope)
}

func buildCommitList(commits []quoad.Commit, remote string) []slack.Block {
	builder := strings.Builder{}
	blocks := []slack.Block{}

	if len(commits) > 15 {
		commits = commits[:15]
		blocks = append(
			blocks,
			slack.NewContextBlock(
				"",
				&slack.TextBlockObject{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"Only last 15 commits shown. *Full changelog <%s|here>*",
						remote,
					),
				},
			),
		)
	}

	for _, commit := range commits {
		smallHash := commit.Hash.String()[:8]

		builder.WriteString(
			fmt.Sprintf(
				"`%s` <%s/commit/%s|%s%s>",
				smallHash,
				remote,
				smallHash,
				addScope(commit.Scope),
				commit.Heading,
			),
		)

		refLinks := []string{}

		for _, ref := range commit.Issues {
			refLinks = append(
				refLinks,
				fmt.Sprintf(
					"<%s/issues/%d|#%d>",
					remote,
					ref,
					ref,
				),
			)
		}

		if len(refLinks) > 0 {
			builder.WriteString(fmt.Sprintf(" _ref %s_", strings.Join(refLinks, ",")))
		}

		builder.WriteString("\r\n")
	}

	blocks = append(
		[]slack.Block{
			slack.SectionBlock{
				Type: "section",
				Text: &slack.TextBlockObject{
					Type: "mrkdwn",
					Text: builder.String(),
				},
			},
		},
		blocks...,
	)

	return blocks
}
