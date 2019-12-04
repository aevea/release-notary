package text

import (
	"strings"

	"github.com/outillage/quoad"
)

func (r *ReleaseNotes) buildCommitLog(commits []quoad.Commit, open bool) string {
	builder := strings.Builder{}

	for index, commit := range commits {
		last := false
		if index+1 == len(commits) {
			last = true
		}
		builder.WriteString(r.buildSingleCommit(commit, last, open))
	}

	builder.WriteString("\n")

	return builder.String()
}
