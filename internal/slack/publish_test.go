package slack

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aevea/quoad"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	t.Skip()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/webhook", req.URL.String())
		assert.Equal(t, "application/json", req.Header["Content-Type"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody, err := ioutil.ReadFile("./testdata/expected_output.json")

		assert.Equal(t, string(body), string(expectedBody))

		_, err = rw.Write([]byte(`ok`))

		assert.NoError(t, err)
	}))

	defer server.Close()

	slack := &Slack{
		WebHookURL: server.URL + "/webhook",
	}

	testData := map[string][]quoad.Commit{
		"features": {
			{Category: "feat", Scope: "ci", Heading: "ci test"},
		},
		"bugs": {
			{Category: "bug", Scope: "", Heading: "huge bug"},
			{Category: "fix", Scope: "", Heading: "bug fix"},
		},
		"chores": {
			{Category: "chore", Scope: "", Heading: "testing", Issues: []int{1, 2}},
			{Category: "improvement", Scope: "", Heading: "this should end up in chores", Issues: []int{3}},
		},
		"others": {
			{Category: "other", Scope: "", Heading: "merge master in something"},
			{Category: "bs", Scope: "", Heading: "random"},
		},
	}

	remote := MockRemote{}

	err := slack.Publish(testData, remote)

	assert.NoError(t, err)
}
