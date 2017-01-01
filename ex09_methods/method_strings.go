package main

import (
	"strconv"
	"fmt"
)

type TwinNumber struct {
	x int
	y int
}

func (tn *TwinNumber) String() string {
	return "(" + strconv.Itoa(tn.x) + " - " + strconv.Itoa(tn.y) + ")"
}

func main()  {
	twiny := &TwinNumber{18, 28}
	fmt.Printf("twiny is : %v \n", twiny) // twiny is : (18 - 28)
	fmt.Println("twiny is :", twiny) // twiny is : (18 - 28)
	fmt.Printf("twiny is: %T \n", twiny) // twiny is: *main.TwinNumber
	fmt.Printf("twiny is: %#v \n", twiny) // twiny is &main.TwinNumber{x:18, y:28}
}
