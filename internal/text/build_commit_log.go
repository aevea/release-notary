package text

import "strings"

func (r *ReleaseNotes) buildCommitLog(commits []Commit) string {
	builder := strings.Builder{}

	for _, commit := range commits {
		builder.WriteString(r.buildSingleCommit(commit))
	}

	builder.WriteString("\n")

	return builder.String()
}
