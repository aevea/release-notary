package internal

// SectionInfo contains the details that we store for each section
type SectionInfo struct {
	ID    string
	Title string
	Icon  string
	// whether the section should not be collapsed by default
	OpenByDefault bool
	Position      int
}

// PredefinedSections contains a list of the predefined sections we consider when generating the release notes
var PredefinedSections = map[string]SectionInfo{
	"features": {Title: "Features", Icon: "rocket", OpenByDefault: true, Position: 0},
	"bugs":     {Title: "Bug fixes", Icon: "bug", Position: 1},
	"chores":   {Title: "Chores and Improvements", Icon: "wrench", Position: 2},
	"breaking": {Title: "Breaking Changes", Icon: "warning", Position: 3},
	"others":   {Title: "Other", Icon: "package", Position: 4},
}

// SortedSections returns a slice of Sections sorted by position
var SortedSections = make([]SectionInfo, len(PredefinedSections))

func init() {
	for identifier, section := range PredefinedSections {
		section.ID = identifier
		SortedSections[section.Position] = section
	}
}
