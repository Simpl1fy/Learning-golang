// to make a file executable, the file should have main function
// and be a part of main package
// if you want a different package, you must use sub-folders
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// datatypes and variable
	var number int = 1
	var name string = "Gourab"
	var boolean bool = true
	var decimal float64 = 1.2
	const pi = 3.142

	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(boolean)
	fmt.Println(decimal)
	fmt.Println(pi)

	// shorthand way of declaring var
	amount := 10
	fmt.Println(amount)
}
