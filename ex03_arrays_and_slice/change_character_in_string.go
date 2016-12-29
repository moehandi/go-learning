package main

import "fmt"

func main() {
	// string in fact is array of bytes
	str := "gopher" // string is immutable
	// to change string we must store it as array of bytes and then convert it back to string
	c := []byte(str)
	c[1] = 'a' // now change o to a
	fmt.Println(string(c)) // now the string is gapher
}
