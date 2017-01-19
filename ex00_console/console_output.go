package main

import "fmt"

// a brief introduction how to output various formatting value to console
// we'll explore package fmt
func main()  {
	// lets define some variable
	var aString string = "im string variable"
	var aNumber int = 23
	var aFloat float64 = 28.9839843
	var aBool bool = true

	// print out to the console
	fmt.Println("fmt.Println return:", aString, aNumber, aFloat, aBool)

	// print out to the console with verb format
	// since fmt.Printf doesn't out new line we must provide it
	fmt.Printf("%v, %v, %v, %v\n", aString, aNumber, aFloat, aBool)

	// lets returning code above to new variable
	// it wont return the concat of string code above but return it lenght of character
	lengthPrintA, err := fmt.Printf("%v, %v, %v, %v\n", aString, aNumber, aFloat, aBool)
	//since it return an error we must check it first
	if err != nil {
		fmt.Println("there is an error")
	}
	fmt.Println(lengthPrintA)

	// with fmt.Sprintf it will return it's concatenation
	// of output to another variable called SprintfResult
	SprintfResult := fmt.Sprintf("Sprintf:%v, %v, %v, %v\n", aString, aNumber, aFloat, aBool)
	fmt.Println(SprintfResult)

	// now we can ignore err code above just replace with _, it means we ignore and don't use the value returning to it
	lengthPrintB, _ := fmt.Printf("%v, %v, %v, %v\n", aString, aNumber, aFloat, aBool)
	fmt.Println(lengthPrintB)

	// lets use Printf to show it's type
	fmt.Printf("we got the type:%T, %T, %T, %T\n", aString, aNumber, aFloat, aBool)
}
