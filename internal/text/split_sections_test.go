package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSections(t *testing.T) {
	dataset := []Commit{
		Commit{Category: "chore", Scope: "", Heading: "testing"},
		Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		Commit{Category: "other", Scope: "", Heading: "merge master in something"},
		Commit{Category: "bs", Scope: "", Heading: "random"},
		Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"},
		Commit{Category: "bug", Scope: "", Heading: "huge bug"},
		Commit{Category: "fix", Scope: "", Heading: "bug fix"},
	}

	expected := Sections{
		Features: []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		Chores:   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		Bugs:     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug"}, Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		Others:   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	sections := SplitSections(dataset)

	assert.Equal(t, expected.Features, sections.Features)
	assert.Equal(t, expected.Chores, sections.Chores)
	assert.Equal(t, expected.Bugs, sections.Bugs)
	assert.Equal(t, expected.Others, sections.Others)
}
