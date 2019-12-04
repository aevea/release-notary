package text

import (
	"log"

	"github.com/outillage/quoad"
)

// SplitSections accepts categorised commits and further organises them into separate sections for release notes
func SplitSections(categorisedCommits []quoad.Commit) map[string][]quoad.Commit {
	categoryMappings := map[string]string{
		"feat":        "features",
		"chore":       "chores",
		"improvement": "chores",
		"bug":         "bugs",
		"other":       "other",
		"fix":         "bugs",
	}

	sections := make(map[string][]quoad.Commit)

	for _, commit := range categorisedCommits {
		var category = categoryMappings[commit.Category]
		if category != "other" {
			log.Println(category)
			sections[category] = append(sections[category], commit)
		}

		if category == "other" || category == "" {
			sections["others"] = append(sections["others"], commit)
		}
	}

	return sections
}
