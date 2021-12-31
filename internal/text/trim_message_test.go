package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimMessage(t *testing.T) {
	testset := map[string]string{
		"ci: sometest \n": "ci: sometest",
		"ci: sometest\n":  "ci: sometest",
		"chore: some github message\n some more text here": "chore: some github message",
		"chore: someother thing":                           "chore: someother thing",
	}

	for input, expected := range testset {
		assert.Equal(t, expected, TrimMessage(input))
	}

}
