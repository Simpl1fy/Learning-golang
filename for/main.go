package main

import "fmt"

func main() {
	fmt.Println("learning for loops")

	// simple for loop similar to c, c++, java
	for i:=0; i<5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// using for loop as while loop
	counter := 0
	for {
		fmt.Printf("%d ", counter)
		if counter == 4 {
			break
		}
		counter++
	}
	fmt.Println()

	// looping through an array
	numbers := []int {1,2,3,4,5}

	for index, value := range numbers {
		fmt.Printf("Index is %d, value is %d\n", index, value)
	}

	fmt.Println()

	// looping through string using range
	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("Index %d, char is %c\n", index, char)
	}
}