package history

import (
	"testing"

	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func setupRepo() *git.Repository {
	repo, _ := git.Init(memory.NewStorage(), memfs.New())

	return repo
}

func createCommit(repo *git.Repository, message string) *object.Commit {
	w, _ := repo.Worktree()

	directory := "./tmp"

	filename := filepath.Join(directory, "example-git-file")
	ioErr := ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	if ioErr != nil {
		panic(ioErr)
	}

	_, addErr := w.Add("example-git-file")
	if addErr != nil {
		panic(addErr)
	}

	commit, _ := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	obj, _ := repo.CommitObject(commit)

	return obj
}

func TestCommitDate(t *testing.T) {
	repo := setupRepo()
	commit := createCommit(repo, "example commit")

	message := CommitMessage(repo, commit.Hash)

	assert.Equal(t, message, "example commit")
}
