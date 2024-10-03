package main

import "fmt"

func main() {
	day := 3

	// basic switch case, no break required
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown")
	}

	// multiple conditions in cases
	month := "January"
	switch month {
	case "January", "December", "November":
		fmt.Println("Winter")
	case "August", "September":
		fmt.Println("Monsoon")
	default:
		fmt.Println("Unknown Weather")
	}

	// evaluation in cases
	temp := 40
	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp > 0 && temp <= 10:
		fmt.Println("Cold")
	case temp > 10 && temp <= 20:
		fmt.Println("Nice")
	case temp > 20 && temp <= 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
