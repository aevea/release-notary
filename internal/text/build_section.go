package text

import "strings"

func (r *ReleaseNotes) buildSection(category string, commits []Commit) string {
	builder := strings.Builder{}

	builder.WriteString(r.buildHeading(category))
	builder.WriteString(r.buildCommitLog(commits, sectionHeadings[category].openByDefault))

	return builder.String()
}
