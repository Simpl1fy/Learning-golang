/*
	The time using in golang is in form 2006-01-02 15:04:05
	where,
	2006 is the year
	01 is the month
	02 is the day
	15 is the hour in 24 hour format
	04 is the minute
	05 is the second
	and can also be written in 12 hour format

	Go will not understand dd-mm-yyyy, to format the time we need to use 2006-01-02, 15:04:05
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time is :", currentTime)
	fmt.Printf("Type of datatype of Time is: %T \n", currentTime)


	// Formatting the time into a readable format
	/*
		Using the golang launch date - 2006-01-02, 15:04:05
	*/
	formatted := currentTime.Format("02-01-2006, 15:04:05")
	fmt.Println("Formatted Date is :", formatted)

	// showing the time in am or pm
	formattedTime := currentTime.Format("3:04:05 PM")
	fmt.Println("Formatted Time is :", formattedTime)


	// Parsing a string to a time format

	// let us take a string of date
	date := "02-11-2024"

	// to parse a string to a date we use time.Parse() function

	layout := "02-01-2006"

	/*
		The layout should be in the format of the date string, in the golang form, then only time package will be able to parse, otherwise logical error.
	*/

	convertedDate, _ := time.Parse(layout, date)
	fmt.Println("The converted string is:", convertedDate)


	// Adding time to date
	new_date := currentTime.Add(48 * time.Hour)
	formatted_new_date := new_date.Format("02/01/2006, Monday")
	fmt.Println("Formatted Date after adding time is:", formatted_new_date)
}
