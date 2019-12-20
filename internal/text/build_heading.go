package text

import (
	"fmt"
	"strings"
)

func (r *ReleaseNotes) buildHeading(category string) string {
	builder := strings.Builder{}

	builder.WriteString("## ")

	icon := fmt.Sprintf(":%v: ", sectionHeadings[category].icon)
	builder.WriteString(icon)

	heading := fmt.Sprintf("%v", sectionHeadings[category].title)

	builder.WriteString(heading)

	builder.WriteString("\n\n")

	return builder.String()
}
