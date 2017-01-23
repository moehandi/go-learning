package main

import (
	"fmt"
	"strings"
)

type Orango struct {
	Firstname string
	Lastname string
}

func UpperCase(o *Orango) {
	o.Firstname = strings.ToUpper(o.Firstname)
	o.Lastname = strings.ToUpper(o.Lastname)
}

func main()  {
	// (1) struct as Value Type
	var o1 Orango
	o1.Firstname = "fisto1"
	o1.Lastname = "lasto1"
	UpperCase(&o1) // because o1 isn't pointer we pass with &
	fmt.Printf("(1) struct => %v, type=> %T \n", o1, o1)
	fmt.Printf("o1, %v %v \n\n", o1.Firstname, o1.Lastname)

	// (2) struct as pointer Type
	o2 := new(Orango)
	o2.Firstname = "firstO2"
	o2.Lastname = "lastO2"
	UpperCase(o2)
	fmt.Printf("(2) struct => %v, type=> %T \n", o2, o2)
	fmt.Printf("o2, %v %v \n\n", o2.Firstname, o2.Lastname)

	// (3) struct as a literal
	o3 := &Orango{"firsto3","lasto3"} // literal use &
	UpperCase(o3) // now argument still point to pointer
	fmt.Printf("(3) struct => %v, type=> %T \n", o3, o3)
	fmt.Printf("o3, %v %v \n\n", o3.Firstname, o3.Lastname)

}
