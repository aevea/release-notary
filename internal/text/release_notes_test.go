package text

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseNotes(t *testing.T) {
	notes := ReleaseNotes{}
	file, err := os.Open("../../expected-output.md")

	assert.NoError(t, err)

	defer file.Close()

	b, err := ioutil.ReadAll(file)

	assert.NoError(t, err)

	expected := string(b)

	sections := map[string][]Commit{
		"features": []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test", Body: "- Body"}},
		"chores":   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing", Body: "- Body"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug", Body: "Body"}},
		"others":   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := notes.Generate(sections)

	assert.Equal(t, expected+"\n", "test heading"+releaseNotes)
}

func TestReleaseNotesSimple(t *testing.T) {
	notes := ReleaseNotes{}
	file, err := os.Open("../../expected-output-simple.md")

	assert.NoError(t, err)

	defer file.Close()

	b, err := ioutil.ReadAll(file)

	assert.NoError(t, err)

	expected := string(b)

	sections := map[string][]Commit{
		"features": []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug"}},
		"others":   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := notes.Generate(sections)

	assert.Equal(t, expected+"\n", "test heading"+releaseNotes)
}

func TestReleaseNotesWithMissingSections(t *testing.T) {
	notes := ReleaseNotes{}
	expected := "\n\n## Features :rocket:\n\n0000000 ci test\n\n"

	sections := map[string][]Commit{
		"features": []Commit{Commit{Heading: "ci test"}},
	}

	releaseNotes := notes.Generate(sections)

	assert.Equal(t, expected, releaseNotes)
}
