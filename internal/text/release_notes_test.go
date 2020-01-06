package text

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/outillage/quoad"
	"github.com/stretchr/testify/assert"
)

func TestReleaseNotes(t *testing.T) {
	notes := ReleaseNotes{Complex: true}
	file, err := os.Open("../../expected-output.md")

	assert.NoError(t, err)

	defer file.Close()

	b, err := ioutil.ReadAll(file)

	assert.NoError(t, err)

	expected := string(b)

	sections := map[string][]quoad.Commit{
		"features": []quoad.Commit{quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test", Body: "- Body"}},
		"chores":   []quoad.Commit{quoad.Commit{Category: "chore", Scope: "", Heading: "testing", Body: "- Body"}, quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores", Issues: []int{12}}},
		"bugs":     []quoad.Commit{quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug", Body: "Body"}},
		"others":   []quoad.Commit{quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, quoad.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := notes.Generate(sections, false)

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

	sections := map[string][]quoad.Commit{
		"features": []quoad.Commit{quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []quoad.Commit{quoad.Commit{Category: "chore", Scope: "", Heading: "testing"}, quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores", Issues: []int{12}}},
		"bugs":     []quoad.Commit{quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"}},
		"others":   []quoad.Commit{quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, quoad.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := notes.Generate(sections, false)

	assert.Equal(t, expected+"\n", "test heading"+releaseNotes)
}

func TestReleaseNotesWithMissingSections(t *testing.T) {
	notes := ReleaseNotes{}
	expected := "\n\n## :rocket: Features\n\n0000000 ci test\n\n"

	sections := map[string][]quoad.Commit{
		"features": []quoad.Commit{quoad.Commit{Heading: "ci test"}},
	}

	releaseNotes := notes.Generate(sections, false)

	assert.Equal(t, expected, releaseNotes)
}
