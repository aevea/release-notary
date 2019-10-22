package text

import "strings"

func buildCommitLog(commits []Commit) string {
	builder := strings.Builder{}

	for _, commit := range commits {
		builder.WriteString(buildSingleCommit(commit))
	}

	builder.WriteString("\n")

	return builder.String()
}
