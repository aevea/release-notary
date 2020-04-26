package text

import (
	"testing"

	"github.com/aevea/quoad"
	"github.com/stretchr/testify/assert"
)

func TestSplitSections(t *testing.T) {
	dataset := []quoad.Commit{
		quoad.Commit{Category: "chore", Scope: "", Heading: "testing"},
		quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"},
		quoad.Commit{Category: "bs", Scope: "", Heading: "random"},
		quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"},
		quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"},
		quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"},
	}

	expected := map[string][]quoad.Commit{
		"features": []quoad.Commit{quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []quoad.Commit{quoad.Commit{Category: "chore", Scope: "", Heading: "testing"}, quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []quoad.Commit{quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"}, quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   []quoad.Commit{quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, quoad.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	sections := SplitSections(dataset)

	assert.Equal(t, expected["features"], sections["features"])
	assert.Equal(t, expected["chores"], sections["chores"])
	assert.Equal(t, expected["bugs"], sections["bugs"])
	assert.Equal(t, expected["others"], sections["others"])
}
