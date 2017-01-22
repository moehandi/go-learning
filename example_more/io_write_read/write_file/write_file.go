package main

import (
	"os"
	"io"
	"fmt"
	"io/ioutil"
)

func main()  {
	text := "string content to write"

	file, err := os.Create("./file.txt")
	errorCheck(err)
	defer file.Close()

	ln, err := io.WriteString(file, text)
	errorCheck(err)

	fmt.Printf("All done write to file %v  characters\n", ln)

	bytes := []byte(text)
	ioutil.WriteFile("./file_bytes.txt", bytes, 0644)
}

func errorCheck(err error)  {
	if err != nil {
		panic(err)
	}
}