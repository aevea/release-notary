package text

import (
	"testing"

	"github.com/aevea/quoad"
	"github.com/stretchr/testify/assert"
)

func TestSplitSections(t *testing.T) {
	dataset := []quoad.Commit{
		{Category: "chore", Scope: "", Heading: "testing"},
		{Category: "feat", Scope: "ci", Heading: "ci test"},
		{Category: "other", Scope: "", Heading: "merge master in something"},
		{Category: "bs", Scope: "", Heading: "random"},
		{Category: "improvement", Scope: "", Heading: "this should end up in chores"},
		{Category: "bug", Scope: "", Heading: "huge bug"},
		{Category: "fix", Scope: "", Heading: "bug fix"},
	}

	expected := map[string][]quoad.Commit{
		"features": {quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   {quoad.Commit{Category: "chore", Scope: "", Heading: "testing"}, quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     {quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"}, quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   {quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, quoad.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	sections := SplitSections(dataset)

	assert.Equal(t, expected["features"], sections["features"])
	assert.Equal(t, expected["chores"], sections["chores"])
	assert.Equal(t, expected["bugs"], sections["bugs"])
	assert.Equal(t, expected["others"], sections["others"])
}
