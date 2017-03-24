package main

import "fmt"

// This example show difference between array, slices and map in:
// 1. initialization
// 2. literal

func main() {

	// TODO ARRAY INITIALIZATION
	var arr1 [3]int = [3]int{} // array must be initialize with size if not it's behave like a slice
	fmt.Println("init array: ", arr1) // out : [0,0,0] because init value tu 0
	fmt.Println("len array: ", len(arr1))
	// assign array value by index
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println("after assign value: ", arr1) // out : [1,2,3]

	// TODO ARRAY LITERAL
	arr2 := [3]int { 6, 7, 8}
	fmt.Println("array literal: ", arr2)

	// TODO SLICE INITIALIZATION
	var slice1 []int = []int{} // slice not initialize with size. it's size can grow base on underlying array
	fmt.Println("init slice: ", slice1)
	fmt.Println("len slice:", len(slice1))

	// assign slice value
	// before assign slice value we must assign slice to underlying array
	slice1 = arr1[:] // we must do this, if not slice cannot be assign because not tied to underlying array
	fmt.Println("len slice now: ", len(slice1)) // len now like array1
	// assign slice value
	slice1[0] = 2
	slice1[1] = 3
	slice1[2] = 4
	fmt.Println("slice after assign: ", slice1)

	slice2 := [3]int{6, 7, 8}
	fmt.Println("slice literal: ", slice2)

	// TODO MAP INITIALIZATION
	var map1 = map[int]string{} // map initialization with type int for key, and string for value
	map1 = make(map[int]string) // we must do this before assing value map
	fmt.Println("map initialization:", map1) // out: map[]
	fmt.Println("map len:", len(map1))

	// assign value to map
	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"
	fmt.Println("map after assign value:", map1) //map[1:one 2:two 3:three]
	fmt.Println("map len after assing value:", len(map1))

	// map literal behave like this
	map2 := map[int]string{
		6 : "six",
		7 : "seven",
		8 : "eight",
	}
	fmt.Println("map literal: ", map2)

	map3 := make(map[int]string)
	fmt.Println("map literal 2 :", map3)
}
