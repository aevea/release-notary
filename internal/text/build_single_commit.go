package text

import "strings"

func buildSingleCommit(commit Commit) string {
	builder := strings.Builder{}

	builder.WriteString("- ")
	// Short version of hash usable on Github
	builder.WriteString(commit.Hash.String()[:7])
	builder.WriteString(" ")
	builder.WriteString(commit.Heading)
	builder.WriteString("\n")

	return builder.String()
}