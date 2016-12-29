package main

import "fmt"

func main()  {
	arr := [5]int{1, 2, 3, 4, 5}

	// see arr[:] it's slice now, it's more idiomatic than passing array to func
	fmt.Println("arrays to slices now, Sum", SumSlice(arr[:]))
}

func SumSlice(arrays []int) int {
	var result int
	for i := range arrays  {
		result += arrays[i]
	}
	return result
}