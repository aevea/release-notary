package text

import (
	"fmt"
	"strings"
)

func (r *ReleaseNotes) buildHeading(category string) string {
	builder := strings.Builder{}

	builder.WriteString("## ")

	heading := fmt.Sprintf("%v ", sectionHeadings[category].title)

	builder.WriteString(heading)

	icon := fmt.Sprintf(":%v:", sectionHeadings[category].icon)

	builder.WriteString(icon)

	builder.WriteString("\n\n")

	return builder.String()
}
