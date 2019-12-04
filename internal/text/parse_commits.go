package text

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	referenceFormatRegex   = regexp.MustCompile(`Refs:?[^\r\n]*`)
	referenceIDFormatRegex = regexp.MustCompile(`\#([0-9]+)`)
	expectedFormatRegex    = regexp.MustCompile(`(?s)^(?P<category>\S+?)?(?P<scope>\(\S+\))?(?P<breaking>!?)?: (?P<heading>[^\n\r]+)?([\n\r]{2}(?P<body>.*))?`)
)

// GetIssueNumbers converts the matches from the reference regular expression to integers
func GetIssueNumbers(matches []string) []int {
	var issueNumbers []int
	for _, match := range matches {
		for _, refID := range referenceIDFormatRegex.FindAllStringSubmatch(match, -1) {
			issueNumber, err := strconv.Atoi(refID[1])

			if err != nil {
				fmt.Println("couldn't convert reference ID to number")
				continue
			}

			issueNumbers = append(issueNumbers, issueNumber)
		}
	}

	return issueNumbers
}

// ParseCommitMessage creates a slice of Commits that contain information about category and scope parsed from commit message
func ParseCommitMessage(commitMessage string) Commit {
	references := referenceFormatRegex.FindAllString(commitMessage, -1)

	commitMessage = referenceFormatRegex.ReplaceAllString(commitMessage, "")

	match := expectedFormatRegex.FindStringSubmatch(commitMessage)

	if len(match) > 0 {
		result := make(map[string]string)

		for i, name := range expectedFormatRegex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		category := result["category"]
		scope := result["scope"]
		heading := result["heading"]
		body := result["body"]

		scope = strings.Replace(scope, "(", "", 1)
		scope = strings.Replace(scope, ")", "", 1)

		return Commit{Category: category, Heading: heading, Scope: scope, Body: strings.TrimRight(body, "\r\n\t "), Issues: GetIssueNumbers(references)}
	}

	return Commit{Category: "other", Heading: commitMessage, Scope: ""}
}
