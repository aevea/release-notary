package history

import (
	"errors"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var (
	errReachedToCommit = errors.New("reached to commit")
)

// CommitsBetween returns a slice of commit hashes between two commits
func CommitsBetween(repo *git.Repository, from plumbing.Hash, to plumbing.Hash) ([]plumbing.Hash, error) {
	cIter, _ := repo.Log(&git.LogOptions{From: from})

	var commits []plumbing.Hash

	err := cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, c.Hash)
		if c.Hash == to {
			return errReachedToCommit
		}
		return nil
	})

	if err == errReachedToCommit {
		return commits, nil
	}
	if err != nil {
		return commits, err
	}
	return commits, nil
}
