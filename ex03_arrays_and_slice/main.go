package main

import "fmt"

func main() {
	fmt.Println("-------- Start of Arrays Section -----------")
	// declare array type int with size 3
	var arr1 [3]int = [3]int {1, 2, 3}
	fmt.Println("arr1:", len(arr1), arr1)
	fmt.Println("element 1: ", arr1[1])

	// try withour var
	arr2 := [3]int {1, 2, 3}
	fmt.Println("arr2:", len(arr2), arr2)
	fmt.Println("element 1: ", arr2[1])

	// try array size with ellipsis "..."
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr3:", len(arr3), arr3)
	fmt.Println("element 1: ", arr3[1])

	// we can't do this. because different size of array, size 5 assign to size 4
	//arr4 := [4]int{1,2,3,4}
	//arr5 := [5]int{1,2,3,4,5}
	//arr4 = arr5
	//fmt.Println(arr4);

	// example with different type: string,
	// declare array type string with length 5
	var arrStr [5]string = [5]string {"one","two","three","four","five"}
	fmt.Println(len(arrStr), arrStr)

	fmt.Println("--------  End of Arrays Section  -----------")

	fmt.Println("-------- Start OF Slices Section -----------")

	// Slice Like array but with empty size, can be built with existing arrays
	// and slice more flexible than array
	slice1 := []int {3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1[1:3]:", slice1[1:3])
	fmt.Println("slice1[:3] :", slice1[:3])
	fmt.Println("slice1[3:] :", slice1[3:])
	fmt.Println("slice1[:]  :", slice1[:]) // equals to prints all

	// See how slice is flexible growth
	slice1 = append(slice1, 10)
	fmt.Println("append '10' slice1 : ", slice1)

	// slices can be loops over
	for key, value := range slice1{
		fmt.Println(key, value)
	}

	// let's try print character inside string 'Hello'
	hello := "Hello"
	var char []rune
	for _, v := range hello {
		char = append(char, v)
	}
	fmt.Println("char result: ", char) // [72 101 108 108 111]
	fmt.Printf("char result %q: ", char) // [72 101 108 108 111]

	// Now see how slice differ from array again
	var arrD []int = []int {1, 2, 3, 4, 5}
	fmt.Println(len(arrD), arrD)

	// create slice size
	var sliceD = make([]int, 5)
	fmt.Println("sliceD: ", len(sliceD), sliceD)
	sliceD = make([]int, 8) // change slice size to 8
	fmt.Println("now Sslice D is: ", len(sliceD), sliceD)

	// get slice from array
	var arrFrom [4]string = [4]string{"one","two","three","four"}
	sliceTo := arrFrom[0:2] // get element from 0 to 2, exclude 2 that's it 0,1
	fmt.Println("now sliceTo is", sliceTo) // printing [one two]

}
