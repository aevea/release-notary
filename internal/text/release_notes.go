package text

import "strings"

// ReleaseNotes generates the output mentioned in the expected-output.md
func ReleaseNotes(sections map[string][]Commit) string {
	builder := strings.Builder{}
	// Extra lines at the start to make sure formatting starts correctly
	builder.WriteString("\n\n")

	if len(sections["features"]) > 0 {
		builder.WriteString(buildSection("features", sections["features"]))
	}

	if len(sections["bugs"]) > 0 {
		builder.WriteString(buildSection("bugs", sections["bugs"]))
	}

	if len(sections["chores"]) > 0 {
		builder.WriteString(buildSection("chores", sections["chores"]))
	}

	if len(sections["others"]) > 0 {
		builder.WriteString(buildSection("others", sections["others"]))
	}

	return builder.String()
}

func buildSection(category string, commits []Commit) string {
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

func buildCommitLog(commits []Commit) string {
	builder := strings.Builder{}

	for _, commit := range commits {
		builder.WriteString("- ")
		// Short version of hash usable on Github
		builder.WriteString(commit.Hash.String()[:7])
		builder.WriteString(" ")
		builder.WriteString(commit.Heading)
		builder.WriteString("\n")
	}

	builder.WriteString("\n")

	return builder.String()
}
