package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url")
	urlString := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	// fmt.Printf("Type of url is: %T\n", urlString)

	// converting url string into url object
	url_obj, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Error converting into url", err)
		return
	}

	// fmt.Printf("Type of new url is: %T\n", url_obj)

	// There url can be divided into 4 parts
	/*
		scheme - https
		host - example.com/8080
		path - path/to/resource/
		query - key=value
	*/

	fmt.Println("Scheme is:", url_obj.Scheme)
	fmt.Println("Host is:", url_obj.Host)
	fmt.Println("Path is:", url_obj.Path)
	fmt.Println("Query is:", url_obj.RawQuery)

	// chaning the url
	url_obj.Path = "/new"
	url_obj.Host = "youtooz.com"
	url_obj.RawQuery = "crazy"

	newURL := url_obj.String()

	fmt.Println("New Url is :", newURL)
}