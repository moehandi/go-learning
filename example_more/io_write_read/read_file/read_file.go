package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	// the file exist that we want to read
	file := "./file_read.txt"

	content, err := ioutil.ReadFile(file)
	errorCheck(err)

	// but it return the bytes
	fmt.Println("Read from file", content)

	result := string(content)
	fmt.Println("Read from file", result)
}

func errorCheck(err error)  {
	if err != nil {
		panic(err)
	}
}
