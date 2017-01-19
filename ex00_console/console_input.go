package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// brief introduction about getting input from console
// to familiarize when we need input such as package fmt.Scanln
func main()  {
	//var value string
	//fmt.Scanln(&value)
	//fmt.Println("your string is", value)

	// the problem code above is only return first input,
	// if we enter space and another second text it will not show
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string text:")
	str, _ := reader.ReadString('\n')
	fmt.Println("you enter this: ", str)

	// wait, if the input is number, or float, how we can get it's real value
	// the value if we want to put to another process
	fmt.Print("Enter a number text: ")
	str, _ = reader.ReadString('\n')
	// to remove \n from string that cause parse error we must trim it with strings.TrimSpace
	fNumber, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value number is: ", fNumber)
		fmt.Println("value after add with 28: ", add(fNumber))
	}
}

func add(f float64) float64 {
	return f + 28
}
