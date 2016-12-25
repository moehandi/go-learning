package main

import "fmt"

func main() {
	// Anonymous function or unnamed function, lets us denote function in expression
	// Anonymous function
	addFunction := func (data ...int) (int, int)  {
		result := 0
		length := len(data)
		for _, d := range data{
			result += d
		}
		return length, result
	}

	length, sum1 := addFunction(10, 20, 30, 40, 50, 60, 70)
	fmt.Println("Len is:", length, "Result is", sum1)
}
