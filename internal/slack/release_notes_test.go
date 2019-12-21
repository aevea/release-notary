package slack

import (
	"fmt"
	"os"
	"testing"

	"github.com/outillage/integrations"
	"github.com/outillage/quoad"
	"github.com/stretchr/testify/assert"
)

func commit(title string) string {
	return fmt.Sprintf("`00000000` <https://example.com/some/thing/commit/00000000|*%s*>", title)
}

func ref(issue int) string {
	return fmt.Sprintf("<https://example.com/some/thing/issues/%d|#%d>", issue, issue)
}

func TestGenerateReleaseNotes(t *testing.T) {
	os.Clearenv()
	remote := integrations.GitRemote{
		Host:    "example.com",
		Project: "some/thing",
	}

	testData := map[string][]quoad.Commit{
		"features": []quoad.Commit{
			quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		},
		"bugs": []quoad.Commit{
			quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"},
			quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"},
		},
		"chores": []quoad.Commit{
			quoad.Commit{Category: "chore", Scope: "", Heading: "testing", Issues: []int{1, 2}},
			quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores", Issues: []int{3}},
		},
		"others": []quoad.Commit{
			quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"},
			quoad.Commit{Category: "bs", Scope: "", Heading: "random"},
		},
	}

	expectedOutput := WebhookMessage(WebhookMessage{
		Blocks: []Block{
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: ":tada: New release for <https://example.com/some/thing|*some/thing*>",
				},
			},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: ":rocket: *Features*",
				},
			},
			Block{
				Type: "context",
				Elements: []content{
					content{
						Type: "mrkdwn",
						Text: "1 commit referencing 0 issues",
					},
				},
			},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: fmt.Sprintf("%s\r\n", commit("ci test")),
				},
			},
			Block{Type: "divider"},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: ":bug: *Bug fixes*",
				},
			},
			Block{
				Type: "context",
				Elements: []content{
					content{
						Type: "mrkdwn",
						Text: "2 commits referencing 0 issues",
					},
				},
			},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"%s\r\n%s\r\n",
						commit("huge bug"),
						commit("bug fix"),
					),
				},
			},
			Block{Type: "divider"},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: ":wrench: *Chores and Improvements*",
				},
			},
			Block{
				Type: "context",
				Elements: []content{
					content{
						Type: "mrkdwn",
						Text: "2 commits referencing 3 issues",
					},
				},
			},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"%s _ref %s,%s_\r\n%s _ref %s_\r\n",
						commit("testing"),
						ref(1),
						ref(2),
						commit("this should end up in chores"),
						ref(3),
					),
				},
			},
			Block{Type: "divider"},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: ":package: *Other*",
				},
			},
			Block{
				Type: "context",
				Elements: []content{
					content{
						Type: "mrkdwn",
						Text: "2 commits referencing 0 issues",
					},
				},
			},
			Block{
				Type: "section",
				Section: content{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"%s\r\n%s\r\n",
						commit("merge master in something"),
						commit("random"),
					),
				},
			},
			Block{Type: "divider"},
		},
	})

	assert.Equal(t, expectedOutput, GenerateReleaseNotes(testData, remote))
}
