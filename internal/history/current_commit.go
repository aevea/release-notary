package history

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// CurrentCommit gets the commit at HEAD
func CurrentCommit(repo *git.Repository) plumbing.Hash {
	head, _ := repo.Head()

	return head.Hash()
}
