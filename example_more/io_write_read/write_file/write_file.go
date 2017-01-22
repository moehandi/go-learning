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

	if err != nil {
		panic(err)
	}

	defer file.Close()

	ln, err := io.WriteString(file, text)
	if err != nil {
		panic(err)
	}

	fmt.Printf("All done write to file %v  characters\n", ln)

	bytes := []byte(text)
	ioutil.WriteFile("./file_bytes.txt", bytes, 0644)
}
