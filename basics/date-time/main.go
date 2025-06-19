package main

import (
	"fmt"
	"time"
)

// For formatting date and time use combination of '02-01-2006 15:04:05, Monday' accordingly.
// For AM/PM use '3:04:05 PM' instead of '15:04:05'.
// For parsing date like 12/06/2003 use '02/01/2006'.

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	formattedTime := currentTime.Format("Monday, 02-01-2006 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	dateStr := "2023-10-01 12:00:00"
	formattedDateStr, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	fmt.Println("Parsed Time:", formattedDateStr)

	new_date := currentTime.Add(24 * time.Hour)
	 fmt.Println("New Date after 24 hours:", new_date.Format("02/01/2006 3:04:05 PM"))
}
