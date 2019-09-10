package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommitMessage(t *testing.T) {

	tests := map[string]Commit{
		"chore: testing":            Commit{Category: "chore", Scope: "", Heading: "testing"},
		"feat(ci): ci test":         Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		"merge master in something": Commit{Category: "other", Scope: "", Heading: "merge master in something"},
	}

	for test, expected := range tests {
		err := ParseCommitMessage(test)
		assert.Equal(t, expected, err)
	}
}
