package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Learning data conversion")
	/*
		Data conversion in go is very similar to c and c++
	*/
	num := 2
	// converting number to float 64
	var num_float = float64(num)

	fmt.Println("Float Number is :", num_float)
	fmt.Printf("Data type is %T\n", num_float)

	// converting integer to string
	number := 1234
	number_string := strconv.Itoa(number)

	fmt.Println("String is :", number_string)
	fmt.Printf("Data type is %T\n", number_string)

	// conversion of string to integer using strconv library

	num_string := "12345"
	num_int, _ := strconv.Atoi(num_string)
	fmt.Println("Integer Number is :", num_int)
	fmt.Printf("Data type is %T\n", num_int)

}
