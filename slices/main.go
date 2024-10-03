package main

import "fmt"

func main() {
	fmt.Println("Learning slices, dynamic in nature similar to vector in c++")
	// numbers := []int{1,2,3,4,5}
	// numbers = append(numbers, 7, 8, 9, 10)
	// fmt.Println("Numbers slice is :", numbers)
	// fmt.Printf("Type of Numbers is %T\n", numbers)
	// fmt.Println("Lenngth of slice is :", len(numbers))

	// using make function to create a slice

	number := make([]int, 3, 5) // make(datatype, size, capacity)

	/*
		Length = Length is the actual elements in the slice
		Capacity = Capacity is the size of the actual array in the memory
	*/

	/*
	When size exceeds capacity, the capacity doubles to fit the number of elements in slice, for the above given example if length=6, then capacity=10(i.e. 5*2)
	*/
	// appending data to number
	number = append(number, 4, 5, 6)

	fmt.Println("Slice", number)
	fmt.Println("Length", len(number))
	fmt.Println("Capacity", cap(number))


}