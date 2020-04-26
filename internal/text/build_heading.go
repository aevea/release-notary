package text

import (
	"fmt"
	"strings"

	"github.com/aevea/release-notary/internal"
)

func (r *ReleaseNotes) buildHeading(category string) string {
	builder := strings.Builder{}

	builder.WriteString("## ")

	icon := fmt.Sprintf(":%v: ", internal.PredefinedSections[category].Icon)

	builder.WriteString(icon)

	heading := fmt.Sprintf("%v", internal.PredefinedSections[category].Title)

	builder.WriteString(heading)

	builder.WriteString("\n\n")

	return builder.String()
}
