package main

import "fmt"

// Declare global variable in this case, package scope. because start with LOWER CASE Letter
var globalVar = "This global variable in this package"

// Declare global variable that can be accessed from outside because start with UPPER CASE Letter
var GlobalVar = "This global variable can be accessed from outside"

func main()  {
	// Print global Variable we declared above
	fmt.Println(globalVar)

	// declare variable with var
	var numVar int8 = 3
	var strVar string = "this is string"

	fmt.Println(numVar)
	fmt.Println(strVar)

	// declare multiple variable
	var numVar1, numVar2, numVar3 = 1, 2, 3
	fmt.Println(numVar1, numVar2, numVar3)

	// Declare multiple variable can be also like this
	var (
		numVar4 = 4
		numVar5 = 5
		numVar6 = 6
	)

	fmt.Println(numVar4, numVar5, numVar6)

	// variable without var
	// but this only works inside function
	noVar := 1
	noVar1, noVar2, noVar3 := 1, 2, 3 // the sequence variable and value must be same
	fmt.Println(noVar)
	fmt.Println(noVar1, noVar2, noVar3)

	// declare constant, constant declaration is without var keyword
	// Remeber constant declared without keyword "var" like variable
	const constA = "constant A"
	fmt.Println(constA)

	// declare multiple const same as declare variable above
	const (
		const1 = "Constant 1"
		const2 = "Constant 2"
	)

	fmt.Println(const1, const2)
}
