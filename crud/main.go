package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*
This function is used for performing get request using go
- use net/http library to perform Get request
- handle error
- close the connection before the end of the main function to ensure memory leaks do not occur
- use NewDecoder function or read data as json to view the data
*/
func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while fetching from link")
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting response:", res.Status)
		return
	}

	/* This is a very basic way of reading the data from response what if we want to store the data directy
	in a struct data type object

	data, error := io.ReadAll(res.Body)

	if error != nil {
		fmt.Println("Error reading data body")
		return
	}

	fmt.Println("Data from body is:", string(data))
	*/

	/*
		Now we are going to use json.NewDecoder().decode() function to decode the json object
		and store the data in the var of the type struct
	*/

	var todo Todo

	er := json.NewDecoder(res.Body).Decode(&todo)
	if er != nil {
		fmt.Println("Error while decoding:", err)
		return
	}

	fmt.Println("Data is:", todo)
}

/*
This is a function for sending post request using go
- First create a object with data for sending
- convert the struct object into json, as the post request can be done in form "Application/JSON"
- for sending post request, we need json string and a reader
  - use string(jsonData) to convert json []byte into a string
  - use strings.NewReader() to create a Reader

- Send a post request to the url and handle error and show result
- result can be shown in 2 ways
  - status
  - data (using io.ReadAll() and converting the result to string)
*/
func performPostRequest() {
	data := Todo{
		UserId:    1,
		Title:     "Gourab Das",
		Completed: true,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myURI := "https://jsonplaceholder.typicode.com/posts"

	res, er := http.Post(myURI, "Application/JSON", jsonReader)
	if er != nil {
		fmt.Println("Error during request:", er)
		return
	}
	defer res.Body.Close()

	// Showing the status code
	// fmt.Println("Status Code:", res.Status)

	// Showing the data recieved
	// we are using io.ReadAll to convert the data into json and show output
	resBody, _ := io.ReadAll(res.Body)
	fmt.Println("data is:", string(resBody))
}

/*
This is a functin for sending put request using go
*/
func performUpdateRequest() {
	data := Todo{
		UserId:    1,
		Title:     "Gourab Das learn go",
		Completed: false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myURI := "https://jsonplaceholder.typicode.com/posts/1"

	req, er := http.NewRequest(http.MethodPut, myURI, jsonReader)
	if er != nil {
		fmt.Println("Error is:", er)
		return
	}
	req.Header.Set("Content-type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error is:", er)
		return
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	fmt.Println("data is:", string(resBody))
}

func main() {
	fmt.Println("CRUD")
	// performGetRequest()
	// performPostRequest()
	performUpdateRequest()
}
