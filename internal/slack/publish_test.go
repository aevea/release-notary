package slack

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/outillage/release-notary/internal/text"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/webhook", req.URL.String())
		assert.Equal(t, "application/json", req.Header["Content-Type"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody := "{\"text\":\"*Features*\\r\\nci test\\r\\n\\r\\n*Bug fixes*\\r\\nhuge bug\\r\\nbug fix\\r\\n\\r\\n*Chores and Improvements*\\r\\ntesting\\r\\nthis should end up in chores\\r\\n\\r\\n*Other*\\r\\nmerge master in something\\r\\nrandom\\r\\n\\r\\n\"}"

		assert.Equal(t, expectedBody, string(body))

		_, err = rw.Write([]byte(`ok`))

		assert.NoError(t, err)
	}))

	defer server.Close()

	slack := &Slack{
		WebHookURL: server.URL + "/webhook",
	}

	testData := map[string][]text.Commit{
		"features": []text.Commit{text.Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		"chores":   []text.Commit{text.Commit{Category: "chore", Scope: "", Heading: "testing"}, text.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		"bugs":     []text.Commit{text.Commit{Category: "bug", Scope: "", Heading: "huge bug"}, text.Commit{Category: "fix", Scope: "", Heading: "bug fix"}},
		"others":   []text.Commit{text.Commit{Category: "other", Scope: "", Heading: "merge master in something"}, text.Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	err := slack.Publish(testData)

	assert.NoError(t, err)
}
