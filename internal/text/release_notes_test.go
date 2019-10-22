package text

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseNotes(t *testing.T) {
	file, err := os.Open("../../expected-output.md")

	defer file.Close()

	assert.NoError(t, err)

	b, err := ioutil.ReadAll(file)

	expected := string(b)

	sections := Sections{
		Features: []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		Chores:   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		Bugs:     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug"}},
		Others:   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := ReleaseNotes(sections)

	assert.Equal(t, expected, releaseNotes)
}

func TestReleaseNotesWithMissingSections(t *testing.T) {
	expected := "\n\n## Features :rocket:\n\n- 0000000 ci test\n\n"

	sections := Sections{
		Features: []Commit{Commit{Heading: "ci test"}},
	}

	releaseNotes := ReleaseNotes(sections)

	assert.Equal(t, expected, releaseNotes)
}
