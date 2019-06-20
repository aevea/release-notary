package history

import (
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// commitDate gets the commit at hash and returns the time of the commit
func commitDate(repo *git.Repository, commit plumbing.Hash) time.Time {
	commitObject, _ := repo.CommitObject(commit)

	when := commitObject.Author.When

	return when
}
