package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero error")
	}
	return a / b, nil
}

// division by 0, always results in +Inf, so to handle error, we can use the error return datatype

func main() {
	fmt.Println("This is the starting of error handling")
	result, _ := divide(10, 0)
	// here we are using underscore operator(read only variable) to store value returned from a function that is not going to be needed
	fmt.Println("Result of divide is", result)
}
