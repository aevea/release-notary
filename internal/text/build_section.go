package text

import (
	"strings"

	"github.com/outillage/quoad"
)

func (r *ReleaseNotes) buildSection(category string, commits []quoad.Commit) string {
	builder := strings.Builder{}

	builder.WriteString(r.buildHeading(category))
	builder.WriteString(r.buildCommitLog(commits, sectionHeadings[category].openByDefault))

	return builder.String()
}
