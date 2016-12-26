package main

import "fmt"

func main()  {
	// operate on no return value
	var java = "java"
	fmt.Println(java)
	operate(&java, "Golang") // will change java to Golang now
	fmt.Println(java)

	// operate on single return value
	x, y := 3, 5
	z := addition(x, y) // assign return value to z
	fmt.Println("x + y =", z)
	fmt.Println("x + y =", addition(x, y) + 1)

	// operate on multiple return values
	data := []int{1,2,3,4,5}
	el1, el2 := getTwoElement(2, 3, data)
	fmt.Println(el1, el2)

	// Variadic function lets us input sequential
	var wordResult = addWord("I'm","Golang","Developer","now")
	fmt.Println(wordResult)

}

// function decalaration no return values & passing pointer as argument
func operate(str *string, replaceWith string) {
	*str = replaceWith // assign new value to address of str with *
}

// function declaration return single int value
func addition(x, y int) int {
	return x + y
}

// function declaration multiple return  int values
func getTwoElement(x, y int, data[] int) (int, int)  {
	return data[x], data[y]
}

// Variadic function with '...'
func addWord(wordList ...string) string {
	var words string
	for _, v := range wordList {
		words = words +" "+ v
	}
	return words
}

