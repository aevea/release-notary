package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aevea/release-notary/internal/release"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	id := 1
	newReleaseNotes := "test"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		url := fmt.Sprintf("/owner/repo/releases/%v", id)
		assert.Equal(t, url, req.URL.String())
		assert.Equal(t, "token testtoken", req.Header["Authorization"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody := fmt.Sprintf("{\"id\":%v,\"tag_name\":\"v1.0.0\",\"name\":\"v1.0.0\",\"body\":\"%v\"}", id, newReleaseNotes)

		assert.Equal(t, expectedBody, string(body))

		_, err = rw.Write([]byte("ok"))

		assert.NoError(t, err)
	}))

	defer server.Close()

	service := &Github{
		APIURL: server.URL,
		client: createClient("testtoken"),
		owner:  "owner",
		repo:   "repo",
	}

	err := service.Publish(&release.Release{ID: id, Tag: "v1.0.0", Name: "v1.0.0", Message: "Description of the release", ReleaseNotes: newReleaseNotes})

	assert.NoError(t, err)
}
