package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("Learning json")
	person := Person{Name: "Gourab Das", Age: 20, IsAdult: true}
	fmt.Println("person object is :", person)

	// convert person object into a json by encoding (Marshalling)
	res, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}

	fmt.Println("Person object in json data is:", string(res))

	// Decoding or unmarshalling a json object

	var PersonData Person
	err = json.Unmarshal(res, &PersonData)

	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}

	fmt.Println("After decoding data is:", PersonData)
}
