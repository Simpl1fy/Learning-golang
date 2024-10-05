package main

import (
	"fmt"
	"os"
)

/*
	In go we can use the os package similar to python to create files, and used the io
	package for writing and reading
*/

func main() {

	/*

		Creating file and writing in the file using function

		fmt.Println("File Handling in Go")
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file", err)
		}
		defer file.Close()
		fmt.Println("file created successfully")

		content := "Hello World"
		_, error := io.WriteString(file, content)

		if error != nil {
			fmt.Println("There was error while writing the file", error)
		}
		fmt.Println("Write in file successful")
	*/

	// Now opening a file using buffer

	/*
		Reading a file, with a buffer based on the size. The file is read in 1024 bytes as mentioned in the code

		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		defer file.Close()

		// creating a buffer to read file content
		buffer := make([]byte, 1024)

		for {	// infinite loop
			n, err := file.Read(buffer)
			if err == io.EOF {		// after reaching eof, the break statement runs, which
				break				// exits the for loop
			}
			if err != nil {			// Error handling
				fmt.Println("Error while reading file:", err)
				return
			}
			fmt.Println(string(buffer[:n]))	// Printing the buffer from 0 to n
		}
	*/

	// reading file using function rather than buffer
	// function is os.ReadFile()

	/*
		If os.ReadFile is used, it takes the entire data into the memory therefore,
		if the file is huge, it may lead to memory leaks. So for huge data, it's
		better to use buffer(byte slice)
	*/

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
	}
	fmt.Println(string(content))
}
