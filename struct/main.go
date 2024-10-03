package main

import "fmt"

/*
	struct is a custom data type similar to c
*/

/*
	We can create a complex struct consisting of different structs
*/

// way to create a struct
type Person struct {
	FirstName string
	LastName string
	Age int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Home string
	Street string
	Pin int
}

// complex struct
type Employee struct {
	Personal_Information Person
	Contact_Details Contact
	Address_Information Address
}

func main() {
	fmt.Println("Learning struct")
	var gourab Person

	fmt.Println("gourab object before initializing:",gourab) // Default value for string is " " and int is 0
	gourab.FirstName = "Gourab"
	gourab.LastName = "Das"
	gourab.Age = 20
	fmt.Println("gourab object after initializing:",gourab)

	// Directly initializing a variable

	person := Person{
		FirstName: "Arijit",
		LastName: "Paul",
		Age: 21,
	}

	fmt.Println("person variable is:", person)

	// using new keyword to make a variable
	person1 := new(Person)
	person1.FirstName = "Arkajit"
	person1.LastName = "Paul"
	person1.Age = 22

	fmt.Println("person1 variable is:", person1)


	// declaring a complex struct to a var
	var emp Employee

	emp.Personal_Information = Person {
		FirstName: "Biraj",
		LastName: "Biswas",
		Age: 24,
	}

	emp.Contact_Details.Email = "biraj@gmail.com"
	emp.Contact_Details.Phone = "56498456465"

	emp.Address_Information = Address{
		Home: "home",
		Street: "bh 34",
		Pin: 718142,
	}
	fmt.Println("Emp1:", emp)
}