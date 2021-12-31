package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type urlTestStruct struct {
	url    string
	commit string
}

func TestLinkToCommit(t *testing.T) {
	tests := map[string]urlTestStruct{
		"https://github.com/commisar-app/commitsar/commit/12345": {url: "https://github.com/commisar-app/commitsar", commit: "12345"},
	}

	for expected, test := range tests {
		link := LinkToCommit(test.url, test.commit)
		assert.Equal(t, expected, link)
	}
}
