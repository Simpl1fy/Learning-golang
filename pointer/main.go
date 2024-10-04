package main

import "fmt"

func call_by_reference(a *int) {
	*a = *a * 2
}

func main() {
	fmt.Println("Learning Pointers")

	/*
		Pointers in go are similar to C, c++, where the pointer is a variable which stores a address of a datatype(int, string, float32, float64). And we can use *(Dereference Operator) to acess the data inside the address
	*/

	num := 10
	ptr := &num

	fmt.Println("Pointer variable is storing:", ptr)
	fmt.Println("Data stored in pointer variable:", *ptr) // Dereferencing operator
	
	/*
	The manual way of assigning a pointer can be written as
	var ptr *int
	*/
	
	/*
	We can also use pointers in function parameters for call by reference
	*/
	call_by_reference(&num)
	fmt.Println("Data stored in pointer variable:", *ptr) // Dereferencing operator
}
