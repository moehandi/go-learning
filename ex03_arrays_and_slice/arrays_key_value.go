package main

import "fmt"

func main() {
	arr1 := [...]string{ 0:"zero", 1:"one", 2:"two", 3:"three", 4:"four", 5:"five"}

	for i := range arr1{
		fmt.Println("array: ",i, "is", arr1[i])
	}

	// only element with index 0, 2, 3, 5 and the other is set to empty
	arr2 := [...]string{ 0:"zero", 2:"two", 3:"three", 5:"five"}

	for i := range arr1{
		fmt.Println("array2: ",i, "is", arr2[i])
	}
}
