package cmd

import (
	history "github.com/outillage/git/pkg"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func getCommits(repo *history.Git) ([]plumbing.Hash, error) {
	currentCommit, err := repo.CurrentCommit()

	if err != nil {
		return nil, err
	}

	lastTag, err := repo.PreviousTag(currentCommit.Hash)

	if err != nil && err != history.ErrPrevTagNotAvailable {
		return nil, err
	}

	// If previous tag is not available, provide empty hash so that all commits are iterated.
	if err == history.ErrPrevTagNotAvailable {
		lastTag = &history.Tag{}
	}

	currentTag, err := repo.CurrentTag()

	if err == history.ErrCommitNotOnTag {
		return nil, history.ErrCommitNotOnTag
	}
	if err != nil {
		return nil, err
	}

	commits, err := repo.BranchDiffCommits(currentTag.Name, lastTag.Name)

	if err != nil {
		return nil, err
	}

	return commits, nil
}
