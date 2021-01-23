package slack

import (
	"fmt"
	"testing"

	"github.com/aevea/quoad"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func commit(title string) string {
	return fmt.Sprintf("`00000000` <https://example.com/some/thing/commit/00000000|*%s*>", title)
}

func ref(issue int) string {
	return fmt.Sprintf("<https://example.com/some/thing/issues/%d|#%d>", issue, issue)
}

func TestGenerateReleaseNotes(t *testing.T) {
	t.Skip()

	// This test needs to be rewriten for the new go-slack structure https://github.com/aevea/release-notary/issues/238
	remote := MockRemote{}

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

	expectedOutput := []slack.Block{}

	assert.ElementsMatch(t, expectedOutput, GenerateReleaseNotes(testData, remote))
}
