package text

import "strings"

func buildSection(category string, commits []Commit) string {
	builder := strings.Builder{}

	builder.WriteString(buildHeading(category))
	builder.WriteString(buildCommitLog(commits))

	return builder.String()
}
