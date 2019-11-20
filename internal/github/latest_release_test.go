package github

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/commitsar-app/release-notary/internal/release"
	"github.com/stretchr/testify/assert"
)

func TestLatestRelease(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/repos/owner/repo/releases/latest", req.URL.String())
		assert.Equal(t, "token testtoken", req.Header["Authorization"][0])

		testFile, err := ioutil.ReadFile("./testdata/latest_release.json")

		assert.NoError(t, err)

		_, err = rw.Write([]byte(string(testFile)))

		assert.NoError(t, err)
	}))

	defer server.Close()

	service := &Github{
		APIURL: server.URL,
		client: createClient("testtoken"),
		owner:  "owner",
		repo:   "repo",
	}

	latest, err := service.LatestRelease()

	assert.NoError(t, err)
	assert.Equal(t, &release.Release{ID: 1, Tag: "v1.0.0", Name: "v1.0.0", Message: "Description of the release", ReleaseNotes: ""}, latest)
}
