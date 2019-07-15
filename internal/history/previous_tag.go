package history

import (
	"errors"
	"log"
	"sort"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	// ErrPrevTagNotAvailable is returned when no previous tag is found.
	ErrPrevTagNotAvailable = errors.New("previous tag is not available")
)

type tag struct {
	hash plumbing.Hash
	time time.Time
}

// PreviousTag sorts tags based on when their commit happened and returns the one previous
// to the current.
func PreviousTag(repo *git.Repository, currentHash plumbing.Hash, debug bool) (plumbing.Hash, error) {
	tagrefs, _ := repo.Tags()
	defer tagrefs.Close()

	var tagHashes []tag

	err := tagrefs.ForEach(func(t *plumbing.Reference) error {
		tagHashes = append(tagHashes, tag{time: commitDate(repo, t.Hash()), hash: t.Hash()})
		return nil
	})

	if err != nil {
		if debug {
			log.Printf("[ERR] getting previous tag failed: %v", err)
		}
		return currentHash, err
	}

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(repo, tagHashes)

	if debug {
		log.Println("Sorted tag output: ")
		for _, taginstance := range sortedTags {
			log.Printf("hash: %v time: %v", taginstance.hash, taginstance.time.UTC())
		}
	}

	// If there are fewer than two tags assume that the currentCommit is the newest tag
	if len(sortedTags) < 2 {
		if debug {
			log.Println("[ERR] previous tag not available")
		}
		return currentHash, ErrPrevTagNotAvailable
	}

	if sortedTags[0].hash != currentHash {
		if debug {
			log.Println("[ERR] current commit does not have a tag attached, building from this commit")
		}
		return sortedTags[0].hash, nil
	}

	if debug {
		log.Printf("success: previous tag found at %v", sortedTags[1].hash)
	}

	return sortedTags[1].hash, nil
}

// sortTags sorts the tags according to when their parent commit happened.
func sortTags(repo *git.Repository, tags []tag) []tag {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].time.After(tags[j].time)
	})

	return tags
}
