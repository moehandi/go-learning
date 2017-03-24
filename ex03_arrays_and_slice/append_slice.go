package main

import "fmt"

func main()  {
	arr1 := []int{1,2,3,4,5}

	arr1 = appendTo(arr1, 6)
	fmt.Println(arr1)

	arr2 := append(arr1, 7)
	fmt.Println(arr2)
}

func appendTo(data []int, value int) []int {
	data = append(data, value)
	return data
}