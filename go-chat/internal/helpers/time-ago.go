package helpers

import (
	"fmt"
	"time"
)

func TimeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration.Hours() >= 24:
		return t.Format("01-02-2006 3:4 PM")
	case duration.Hours() >= 1:
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	case duration.Minutes() >= 1:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	default:
		seconds := int(duration.Seconds())
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}
