package main

import "fmt"

func hello() {
	fmt.Println("Hello World through a function")
}

// input parameters
func sum(a, b int) int { // func name, input parameters, return data type
	return a + b
}

// named returns
func float_sum(a, b float64) (result float64) {
	result = a + b
	return
}

func main() {
	fmt.Println("We are learning functions.")
	hello()
	res := sum(1, 2)
	fmt.Println("Sum is", res)
	sum := float_sum(5.5, 5.5)
	fmt.Println("Sum is", sum)
}
