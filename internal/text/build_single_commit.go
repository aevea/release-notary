package text

import (
	"fmt"
	"strings"

	"github.com/outillage/quoad"
)

func (r *ReleaseNotes) buildSingleCommit(commit quoad.Commit, isLast, open bool) string {
	builder := strings.Builder{}

	if !r.Complex || commit.Body == "" {
		simpleCommitMessage := buildSimpleCommit(commit)
		builder.WriteString(simpleCommitMessage)
	} else {
		commitMessage := buildFullCommit(commit, open)
		builder.WriteString(commitMessage)
	}

	builder.WriteString("\n")

	return builder.String()
}

func buildSimpleCommit(commit quoad.Commit) string {
	builder := strings.Builder{}

	// Short version of hash usable on Github
	builder.WriteString("- ")
	builder.WriteString(commit.Hash.String()[:7])
	builder.WriteString(" ")

	if commit.Scope != "" {
		builder.WriteString("**(")
		builder.WriteString(commit.Scope)
		builder.WriteString(")**")
		builder.WriteString(" ")
	}

	builder.WriteString(commit.Heading)

	for _, refID := range commit.Issues {
		builder.WriteString(fmt.Sprintf(" #%d", refID))
	}

	return builder.String()
}

func buildFullCommit(commit quoad.Commit, open bool) string {
	builder := strings.Builder{}

	// closed receives empty string
	openString := ""
	if open {
		openString = " open"
	}
	detailsWrapperStart := fmt.Sprintf("<details%v>", openString)
	builder.WriteString(detailsWrapperStart)
	builder.WriteString("<summary>")
	builder.WriteString(commit.Hash.String()[:7])
	builder.WriteString(" ")
	builder.WriteString(commit.Heading)
	builder.WriteString("</summary>")
	builder.WriteString("\n\n")
	builder.WriteString(commit.Body)

	for _, refID := range commit.Issues {
		builder.WriteString(fmt.Sprintf(" #%d", refID))
	}

	builder.WriteString("\n\n")
	builder.WriteString("</details>")

	return builder.String()
}
