package slack

import (
	"testing"

	"github.com/commitsar-app/release-notary/internal/text"
	"github.com/stretchr/testify/assert"
)

func TestGenerateReleaseNotes(t *testing.T) {
	testData := map[string][]text.Commit{
		"features": []text.Commit{text.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []text.Commit{text.Commit{Category: "chore", Scope: "", Heading: "testing"}, text.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []text.Commit{text.Commit{Category: "bug", Scope: "", Heading: "huge bug"}, text.Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   []text.Commit{text.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, text.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	assert.Equal(t, "*Features*\r\nci test\r\n\r\n*Bug fixes*\r\nhuge bug\r\nbug fix\r\n\r\n*Chores and Improvements*\r\ntesting\r\nthis should end up in chores\r\n\r\n*Other*\r\nmerge master in something\r\nrandom\r\n\r\n", GenerateReleaseNotes(testData))
}
