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

	expected := map[string][]Commit{
		"features": []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug"}, Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	sections := SplitSections(dataset)

	assert.Equal(t, expected["features"], sections["features"])
	assert.Equal(t, expected["chores"], sections["chores"])
	assert.Equal(t, expected["bugs"], sections["bugs"])
	assert.Equal(t, expected["others"], sections["others"])
}
