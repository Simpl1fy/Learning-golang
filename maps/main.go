package main

import "fmt"

/*
	Maps is unordered data structure for key value pair, similar to hash maps(c, c++) or dictionaries(pyton)
*/

func main() {
	fmt.Println("Maps in go")

	// making a map, using map function
	students := make(map[string]int)

	students["Gourab"] = 100
	students["Arijit"] = 70
	students["Arkajit"] = 80
	students["Biraj"] = 10
	students["Prince"] = 99
	
	fmt.Println("Marks of Gourab is", students["Gourab"])

	// deleting key value pairs in maps
	delete(students, "Biraj")

	// checking if key exists or not
	marks, exists := students["Biraj"]

	fmt.Println("Biraj exists:", exists)
	fmt.Println("Marks obtained by biraj", marks)

	// looping through maps using for range
	for key, value := range students {
		fmt.Printf("key is %s and value is %d\n", key, value)
	}

	// way to initialize a map during declaration, we do not use make
	/*
		make is used to declare varaibles without initialing any value to them
	*/
	person := map[string] int {
		"alice": 10,
		"bob": 50,
		"prince": 45,
	}

	for key, value := range person {
		fmt.Printf("---------------key is %s and value is %d\n", key, value)
	}
}
