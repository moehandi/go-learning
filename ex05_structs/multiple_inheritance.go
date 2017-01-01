// Multiple inheritance can be implemented by embedding all the necessary parent type
// in the under construction
package main

import "fmt"

type Book struct { }
func (c *Book) Page() string {

	return "Page"
}

type Directory struct { }
func (d *Directory ) Rack() string {
	return "Rack"
}

// multiple inheritance, inherit from Camera and Phone
type BookDirectory struct {
	Book
	Directory
}

func main() {
	bd := new(BookDirectory)
	fmt.Println("Our new BookDirectory exhibits multiple behaviors ...")
	fmt.Println("Got behavior of a Book:", bd.Page())
	fmt.Println("Got behavior of Directory too:", bd.Rack())
}
