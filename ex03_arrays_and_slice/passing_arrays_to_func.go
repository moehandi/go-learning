package main

import "fmt"

func main()  {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrays sum", Sum(arr))
}

// it's not flexible and idiomatic since we must defined same length of arrayas [5]int
// see passing_slices_to_func.go for better work
func Sum(arrays [5]int) int {
	var result int
	for i := range arrays  {
		result += arrays[i]
	}
	return result
}