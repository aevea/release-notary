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
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/webhook", req.URL.String())
		assert.Equal(t, "application/json", req.Header["Content-Type"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody, err := ioutil.ReadFile("./testdata/expected_output.txt")

		type BlockList struct {
			Blocks []Block
		}

		var receivedBlocks, expectedBlocks BlockList

		rbErr, ebErr := json.Unmarshal(body, &receivedBlocks), json.Unmarshal(expectedBody, &expectedBlocks)

		assert.NoError(t, rbErr)
		assert.NoError(t, ebErr)
		assert.ElementsMatch(t, expectedBlocks.Blocks, receivedBlocks.Blocks)

		_, err = rw.Write([]byte(`ok`))

		assert.NoError(t, err)
	}))

	defer server.Close()

	slack := &Slack{
		WebHookURL: server.URL + "/webhook",
	}

	testData := map[string][]quoad.Commit{
		"features": []quoad.Commit{
			quoad.Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		},
		"bugs": []quoad.Commit{
			quoad.Commit{Category: "bug", Scope: "", Heading: "huge bug"},
			quoad.Commit{Category: "fix", Scope: "", Heading: "bug fix"},
		},
		"chores": []quoad.Commit{
			quoad.Commit{Category: "chore", Scope: "", Heading: "testing", Issues: []int{1, 2}},
			quoad.Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores", Issues: []int{3}},
		},
		"others": []quoad.Commit{
			quoad.Commit{Category: "other", Scope: "", Heading: "merge master in something"},
			quoad.Commit{Category: "bs", Scope: "", Heading: "random"},
		},
	}

	remote := MockRemote{}

	err := slack.Publish(testData, remote)

	assert.NoError(t, err)
}
