package history

import (
	"log"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// CurrentCommit gets the commit at HEAD
func CurrentCommit(repo *git.Repository, debug bool) plumbing.Hash {
	head, _ := repo.Head()

	currentCommit := head.Hash()

	if debug {
		log.Println("current commmit: ", currentCommit)
	}

	return currentCommit
}
