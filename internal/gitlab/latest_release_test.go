package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/commitsar-app/release-notary/internal/release"
	"github.com/stretchr/testify/assert"
)

func TestLatestRelease(t *testing.T) {
	projectID := 1
	tagName := "v1.0.0"

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		path := fmt.Sprintf("/projects/%v/repository/tags/%v", projectID, tagName)
		assert.Equal(t, path, req.URL.String())
		assert.Equal(t, "testtoken", req.Header["Private-Token"][0])

		testFile, err := ioutil.ReadFile("./testdata/latest_release.json")

		assert.NoError(t, err)

		_, err = rw.Write([]byte(string(testFile)))

		assert.NoError(t, err)
	}))

	defer server.Close()

	service := &Gitlab{
		apiURL:    server.URL,
		client:    createClient("testtoken"),
		projectID: projectID,
		tagName:   tagName,
	}

	latest, err := service.LatestRelease()

	assert.NoError(t, err)
	assert.Equal(t, &release.Release{ID: 1, Tag: "v1.0.0", Name: "v1.0.0", Message: "Some description", ReleaseNotes: ""}, latest)
}
