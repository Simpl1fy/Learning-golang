package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name: ")
	// var name string

	// // fmt.Scan reads only upto " "
	// fmt.Scan(&name)
	// fmt.Printf("Hello Mr. %s", name)

	// to read full name we need to use package bufio
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("Hello Mr.", name)
}
