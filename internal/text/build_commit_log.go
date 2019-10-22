package text

import "strings"

func (r *ReleaseNotes) buildCommitLog(commits []Commit, open bool) string {
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
