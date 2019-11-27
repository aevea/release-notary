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

	commitObject, err := repo.Commit(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "feat: third file\n", commitObject.Message)

	commitObject, err = repo.Commit(commits[1])
	assert.NoError(t, err)
	assert.Equal(t, "feat: fourth file\n", commitObject.Message)
}

func TestGetCommitsSingleBranch(t *testing.T) {
	repo, err := history.OpenGit("../testdata/single_branch_tags", false)

	assert.NoError(t, err)

	commits, err := getCommits(repo)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(commits))

	commitObject, err := repo.Commit(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "feat: third file\n", commitObject.Message)

	commitObject, err = repo.Commit(commits[1])
	assert.NoError(t, err)
	assert.Equal(t, "feat: second file\n", commitObject.Message)
}
