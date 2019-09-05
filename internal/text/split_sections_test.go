package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSections(t *testing.T) {
	dataset := []CategorisedCommit{
		CategorisedCommit{Category: "chore", Scope: "", Heading: "testing"},
		CategorisedCommit{Category: "feat", Scope: "ci", Heading: "ci test"},
		CategorisedCommit{Category: "other", Scope: "", Heading: "merge master in something"},
		CategorisedCommit{Category: "bs", Scope: "", Heading: "random"},
		CategorisedCommit{Category: "improvement", Scope: "", Heading: "this should end up in chores"},
		CategorisedCommit{Category: "bug", Scope: "", Heading: "huge bug"},
	}

	expected := Sections{
		Features: []string{"ci test"},
		Chores:   []string{"testing", "this should end up in chores"},
		Bugs:     []string{"huge bug"},
		Others:   []string{"merge master in something", "random"},
	}

	sections := splitSections(dataset)

	assert.Equal(t, expected.Features, sections.Features)
	assert.Equal(t, expected.Chores, sections.Chores)
	assert.Equal(t, expected.Bugs, sections.Bugs)
	assert.Equal(t, expected.Others, sections.Others)
}
