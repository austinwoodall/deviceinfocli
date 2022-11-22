package utility

import "fmt"

func TimeConversion(time int) int {
	days := time / (60 * 60 * 24)
	minutes := time / 60
	hours := minutes / 60

	format, _ := fmt.Printf("%d Day(s) %d Hour(s) %d Minute(s) ", days, hours, minutes)
	return format
}
