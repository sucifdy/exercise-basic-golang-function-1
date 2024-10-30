package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	// Array of month names
	monthNames := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	// Format day with leading zero if necessary
	formattedDay := fmt.Sprintf("%02d", day)
	
	// Get the month name (month is 1-indexed)
	monthName := monthNames[month-1]
	
	// Format the final date string
	formattedDate := fmt.Sprintf("%s-%s-%d", formattedDay, monthName, year)
	
	return formattedDate
}

// debugg
func main() {
	fmt.Println(DateFormat(1, 1, 2012))    // Output: "01-January-2012"
	fmt.Println(DateFormat(31, 12, 2020))  // Output: "31-December-2020"
}
