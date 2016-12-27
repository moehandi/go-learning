package main

import "fmt"

func main() {
	n := 0
	result := &n
	add(2, 4, result)
	fmt.Println("new result ", *result)

}

// change value by pointer ( pass by pointer) not pass by value
func add(a, b int, result *int) {
	*result = a * b
}
