/*
	Defer is a keyword, which is used when we want to execute a statement just before
	the ending of execution of program.

	When there are multiple defer keywords, it forms a stack (LIFO), the last defer statement, will execute
	first and so on and so forth.
*/

package main

import "fmt"

func add(a, b int) int {
	return a+b
}

func main() {
	fmt.Println("Starting of the Program")
	data := add(5,6)
	defer fmt.Println("Result of Addition is:", data)
	defer fmt.Println("Middle of the Program")
	fmt.Println("Ending of the Program")
}

/*
	Here the multiple defer keyword creates a stack,
	| fmt.Println("Middle of the Program") 		  |
	| fmt.Println("Result of Addition is:", data) |
	|_____________________________________________|
*/