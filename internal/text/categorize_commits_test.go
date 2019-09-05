package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoriseCommits(t *testing.T) {
	commits := []string{"chore: testing", "feat(ci): ci test", "merge master in something"}

	categorized := CategoriseCommits(commits)

	expected := []CategorisedCommit{
		CategorisedCommit{Category: "chore", Scope: "", Heading: "testing"},
		CategorisedCommit{Category: "feat", Scope: "ci", Heading: "ci test"},
		CategorisedCommit{Category: "other", Scope: "", Heading: "merge master in something"},
	}

	assert.Equal(t, expected, categorized)
}
