package text

import "strings"

// ReleaseNotes generates the output mentioned in the expected-output.md
func ReleaseNotes(sections Sections) string {
	builder := strings.Builder{}
	// Extra lines at the start to make sure formatting starts correctly
	builder.WriteString("\n\n")

	if len(sections.Features) > 0 {
		builder.WriteString(buildSection("features", sections.Features))
	}

	if len(sections.Bugs) > 0 {
		builder.WriteString(buildSection("bugs", sections.Bugs))
	}

	if len(sections.Chores) > 0 {
		builder.WriteString(buildSection("chores", sections.Chores))
	}

	if len(sections.Others) > 0 {
		builder.WriteString(buildSection("others", sections.Others))
	}

	return builder.String()
}

func buildSection(category string, commits []string) string {
	builder := strings.Builder{}

	builder.WriteString(buildHeading(category))
	builder.WriteString(buildCommitLog(commits))

	return builder.String()
}

func buildHeading(category string) string {
	builder := strings.Builder{}

	builder.WriteString("## ")

	if category == "features" {
		builder.WriteString("Features ")
		builder.WriteString(EmoticonFeature)
	}

	if category == "bugs" {
		builder.WriteString("Bug fixes ")
		builder.WriteString(EmoticonBug)
	}

	if category == "chores" {
		builder.WriteString("Chores and Improvements ")
		builder.WriteString(EmoticonChores)
	}

	if category == "others" {
		builder.WriteString("Other ")
		builder.WriteString(EmoticonOthers)
	}

	builder.WriteString("\n\n")

	return builder.String()
}

func buildCommitLog(commits []string) string {
	builder := strings.Builder{}

	for _, commit := range commits {
		builder.WriteString("- ")
		builder.WriteString(commit)
		builder.WriteString("\n")
	}

	builder.WriteString("\n")

	return builder.String()
}
