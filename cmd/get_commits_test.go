package cmd

import (
	"testing"

	history "github.com/outillage/git/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetCommitsDetachedTags(t *testing.T) {
	repo, err := history.OpenGit("../testdata/detached_tag", false)

	assert.NoError(t, err)

	commits, err := getCommits(repo)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(commits))

	tests := map[int]string{
		0: "feat: third file\n",
		1: "feat: fourth file\n",
	}

	for index, message := range tests {
		commitObject, err := repo.Commit(commits[index])
		assert.NoError(t, err)
		assert.Equal(t, message, commitObject.Message)
	}
}

func TestGetCommitsSingleBranch(t *testing.T) {
	repo, err := history.OpenGit("../testdata/single_branch_tags", false)

	assert.NoError(t, err)

	commits, err := getCommits(repo)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(commits))

	tests := map[int]string{
		0: "feat: third file\n",
		1: "feat: second file\n",
	}

	for index, message := range tests {
		commitObject, err := repo.Commit(commits[index])
		assert.NoError(t, err)
		assert.Equal(t, message, commitObject.Message)
	}
}

func TestGetCommitsReleaseCommit(t *testing.T) {
	repo, err := history.OpenGit("../testdata/release_commit", false)

	assert.NoError(t, err)

	commits, err := getCommits(repo)

	assert.NoError(t, err)
	assert.Equal(t, 4, len(commits))

	tests := map[int]string{
		0: "feat: release v0.0.1\n",
		1: "feat: third file\n",
		2: "feat: second file\n",
		3: "feat: first file\n",
	}

	for index, message := range tests {
		commitObject, err := repo.Commit(commits[index])
		assert.NoError(t, err)
		assert.Equal(t, message, commitObject.Message)
	}
}

func TestGetCommitsReleaseNonCompliantCommit(t *testing.T) {
	repo, err := history.OpenGit("../testdata/release_noncompliant_commit", false)

	assert.NoError(t, err)

	commits, err := getCommits(repo)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(commits))

	tests := map[int]string{
		0: "release v0.0.2\n",
		1: "feat: fourth file\n",
	}

	for index, message := range tests {
		commitObject, err := repo.Commit(commits[index])
		assert.NoError(t, err)
		assert.Equal(t, message, commitObject.Message)
	}
}
