package main

import "fmt"

func main()  {

	// REMEMBER !
	//The key type can be any type for which the operations == and != are defined, like string, int, float.
	//So arrays, slices and structs cannot be used as key type, but pointers and interface types can. One
	//way to use structs as a key is to provide them with a Key() or Hash() method, so that a unique
	//numeric or string key can be calculated from the struct’s fields.

	// the integer are mapped to address of function
	mapFunction1 := map[int] func() int {
		1 : func() int {
			return 8
		},
		2 : func() int {
			return 9
		},
		3 : func() int {
			return 10
		},
	}

	fmt.Println("map function1", mapFunction1) // out: map[1:0x47b8a0 2:0x47b8b0 3:0x47b8c0]

	// TODO Another mapFunction Example with string key value
	mapFunction2 := map[string] func() string {
		"one" : func() string {
			return "return one"
		},
		"two" : func() string {
			return "return two"
		},
		"three" : func() string {
			return "return three"
		},
	}
	fmt.Println("map function2: ", mapFunction2)
}