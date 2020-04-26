package cmd

import (
	history "github.com/aevea/git/v2"
	"github.com/go-git/go-git/v5/plumbing"
)

func getCommits(repo *history.Git) ([]plumbing.Hash, error) {
	currentCommit, err := repo.CurrentCommit()

	if err != nil {
		return nil, err
	}

	currentTag, err := repo.CurrentTag()

	if err == history.ErrCommitNotOnTag {
		return nil, history.ErrCommitNotOnTag
	}
	if err != nil {
		return nil, err
	}

	lastTag, err := repo.PreviousTag(currentCommit.Hash)

	if err != nil && err != history.ErrPrevTagNotAvailable {
		return nil, err
	}

	var commits []plumbing.Hash

	// If previous tag is not available, or both tags are the same,
	// provide empty hash so that all commits are iterated.
	if (err == history.ErrPrevTagNotAvailable) || (lastTag == currentTag) {
		lastTag = &history.Tag{}
		commits, err = repo.CommitsBetween(currentTag.Hash, lastTag.Hash)
	} else {
		commits, err = repo.BranchDiffCommits(currentTag.Name, lastTag.Name)
	}

	if err != nil {
		return nil, err
	}

	return commits, nil
}
