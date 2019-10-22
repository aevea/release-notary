package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFullCommit(t *testing.T) {
	testCommit := Commit{Heading: "stuff", Body: "hi"}

	text := buildFullCommit(testCommit, false)

	expectedText := "<details><summary>0000000 stuff</summary>\n\nhi\n\n</details>"
	assert.Equal(t, expectedText, text)

	expectedOpenText := "<details open><summary>0000000 stuff</summary>\n\nhi\n\n</details>"
	openText := buildFullCommit(testCommit, true)
	assert.Equal(t, expectedOpenText, openText)
}
