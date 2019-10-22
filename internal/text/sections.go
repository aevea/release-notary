package text

type sectionInfo struct {
	title string
	icon  string
}

var sectionHeadings = map[string]sectionInfo{

	"breaking": sectionInfo{title: "Breaking Changes", icon: "warning"},
	"features": sectionInfo{title: "Features", icon: "rocket"},
	"chores":   sectionInfo{title: "Chores and Improvements", icon: "wrench"},
	"bugs":     sectionInfo{title: "Bug fixes", icon: "bug"},
	"others":   sectionInfo{title: "Other", icon: "package"},
}
