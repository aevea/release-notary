package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommitMessage(t *testing.T) {

	tests := map[string]Commit{
		"chore: testing\n":                                         Commit{Category: "chore", Scope: "", Heading: "testing"},
		"feat(ci): ci test\n":                                      Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		"merge master in something\n":                              Commit{Category: "other", Scope: "", Heading: "merge master in something\n"},
		"chore: test\n\nsomething more here":                       Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here"},
		"chore: test\n\nsomething more here\nRefs: #12":            Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\n\tRefs: #12":          Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\n\t Refs: #12":         Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\nRefs: #12\nRefs: #13": Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
		"chore: test\n\nsomething more here\nRefs: #12, #13":       Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
		"chore: test\n\nsomething more here\nRefs: #12 and #13":    Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
	}

	for test, expected := range tests {
		err := ParseCommitMessage(test)
		assert.Equal(t, expected, err)
	}
}
