package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseNotes(t *testing.T) {
	expected := "## Features :rocket:\n\n- ci test\n\n## Bug fixes :bug:\n\n- huge bug\n\n## Chores and Improvements :wrench:\n\n- testing\n- this should end up in chores\n\n## Other :package:\n\n- merge master in something\n- random\n\n"

	sections := Sections{
		Features: []string{"ci test"},
		Chores:   []string{"testing", "this should end up in chores"},
		Bugs:     []string{"huge bug"},
		Others:   []string{"merge master in something", "random"},
	}

	releaseNotes := ReleaseNotes(sections)

	assert.Equal(t, expected, releaseNotes)
}
