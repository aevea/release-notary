package internal

// SectionInfo contains the details that we store for each section
type SectionInfo struct {
	Title string
	Icon  string
	// whether the section should not be collapsed by default
	OpenByDefault bool
}

// PredefinedSections contains a list of the predefined sections we consider when generating the release notes
var PredefinedSections = map[string]SectionInfo{
	"breaking": SectionInfo{Title: "Breaking Changes", Icon: "warning"},
	"features": SectionInfo{Title: "Features", Icon: "rocket", OpenByDefault: true},
	"chores":   SectionInfo{Title: "Chores and Improvements", Icon: "wrench"},
	"bugs":     SectionInfo{Title: "Bug fixes", Icon: "bug"},
	"others":   SectionInfo{Title: "Other", Icon: "package"},
}
