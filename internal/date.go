package internal

import "time"

func ParseDate(dateStr string) (time.Time, error) {
	const layout = "1/2/2006 3:04:05 PM"
	return time.Parse(layout, dateStr)
}
