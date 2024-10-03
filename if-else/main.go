package main

import "fmt"

func main() {
	x := 1
	if x > 5 {
		fmt.Printf("%d is greater than 5\n", x)
	} else {
		fmt.Printf("%d is smaller than 5\n", x)
	}

	z := 10
	if z >= 10 {
		fmt.Println("z is greater then or equal to 10")
	} else if z >= 5 {
		fmt.Println("z is smaller then 10 but greater than or equal to 5")
	} else {
		fmt.Println("z is smaller than 5")
	}
}
