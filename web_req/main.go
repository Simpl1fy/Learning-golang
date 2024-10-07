package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	res, err := http.Get("https://dummyjson.com/todos/1")

	
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	
	defer res.Body.Close()  // closing the connection to save memory

	fmt.Printf("Type of response is %T\n", res)

	// reading the response body
	data, error := io.ReadAll(res.Body)

	if error != nil {
		fmt.Println("Error reading response:", error)
		return
	}

	fmt.Println("Data is :", string(data))
}
