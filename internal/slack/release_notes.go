package slack

import (
	"fmt"
	"os"
	"strings"

	"github.com/aevea/quoad"
	"github.com/aevea/release-notary/internal"
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
func GenerateReleaseNotes(sections map[string][]quoad.Commit, remote GitRemoter) WebhookMessage {
	blocks := []Block{buildReleaseTitle(remote)}

	for name, commits := range sections {
		if len(commits) > 0 {
			sectionInfo := internal.PredefinedSections[name]

			sectionTitle := Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: fmt.Sprintf(":%s: *%s*", sectionInfo.Icon, sectionInfo.Title),
				},
			}

			sectionContext := Block{
				Type: "context",
				Elements: []content{
					content{
						Type: "mrkdwn",
						Text: fmt.Sprintf(
							"%s referencing %s",
							pluralize("commit", len(commits)),
							pluralize("issue", countReferences(commits)),
						),
					},
				},
			}

			sectionCommits := Block{
				Type:    "section",
				Section: buildCommitList(commits, remote.GetRemoteURL()),
			}

			blocks = append(
				blocks,
				sectionTitle,
				sectionContext,
				sectionCommits,
				Block{Type: "divider"},
			)
		}
	}

	return WebhookMessage{Blocks: blocks}
}

func buildReleaseTitle(remote GitRemoter) Block {
	// This is also quite hacky, it shouldn't be done here, but somewhere before
	release, isGithub := os.LookupEnv("GITHUB_REPOSITORY")
	remoteURL := remote.GetRemoteURL()

	// TODO: Improve this logic
	if !isGithub {
		return Block{
			Type: "section",
			Section: content{
				Type: "mrkdwn",
				Text: fmt.Sprintf(
					":tada: Release <%s/releases/tag/%s|*%s*> for <%s|*%s*>",
					remoteURL,
					release,
					release,
					remoteURL,
					remote.Project(),
				),
			},
		}
	}

	// For GitHub it will be skipped for now :/ We'll need to fetch it via the API
	return Block{
		Type: "section",
		Section: content{
			Type: "mrkdwn",
			Text: fmt.Sprintf(":tada: New release for <%s|*%s*>", remoteURL, remote.Project()),
		},
	}
}

func buildCommitList(commits []quoad.Commit, remote string) content {
	builder := strings.Builder{}

	for _, commit := range commits {
		smallHash := commit.Hash.String()[:8]

		builder.WriteString(
			fmt.Sprintf(
				"`%s` <%s/commit/%s|*%s*>",
				smallHash,
				remote,
				smallHash,
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

	section := content{Type: "mrkdwn", Text: builder.String()}

	return section
}
