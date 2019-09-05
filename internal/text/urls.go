package text

import (
	"fmt"
)

// LinkToCommit to commit provides a formatted link to a commit
func LinkToCommit(projectURL string, commitID string) string {
	link := fmt.Sprintf("%v/commit/%v", projectURL, commitID)

	return link
}