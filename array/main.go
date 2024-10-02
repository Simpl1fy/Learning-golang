/*
In go, when an a array is initialized or declared, the default value is stored in the array
Example - var arr[5] int - {0 0 0 0 0}

	var arr[5] float64 - {0 0 0 0 0}
	var arr[5] bool - {false false false false false}
	for data type string the default value is ""
	var arr[5] string - {     }
*/
package main

import "fmt"

func main() {
	fmt.Println("Learning Array in Go")

	// declaring a array
	var names [5]string
	names[0] = "gourab"
	names[1] = "arijit"

	fmt.Println("Names stored in array", names)

	// length of array
	fmt.Println("Length of array names is", len(names))

	// initializing an array
	var number = [5]int{1, 2, 3, 4, 5}

	// iterating over arrays, using for loop
	for i := 0; i < len(number); i++ {
		fmt.Printf("%d ", number[i])
	}
	fmt.Printf(" \n")

	// iteration over arrays using range function
	for index, value := range number {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}

	var name [5]string
	name[0] = "gourab"
	name[3] = "prince"
	fmt.Printf("The name array in quoted is %q\n", name)
}
