package main

import (
	"math"
	"fmt"
)

// we named interface Shapes, to avoid redeclared interface name in interface_declaration.go
type Shapes interface {
	Area() float32
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func main()  {
	c1 := new(Circle)
	c1.radius = 7

	r1 := new(Rectangle)
	r1.width = 29
	r1.length = 18

	// add r1, c1 type to arrays of interface type
	// with this we got polymorphism
	shapes := []Shapes{c1, r1}

	for _,v := range shapes {
		fmt.Printf("Shape type: %T, Area(): %v \n", v, v.Area())
	}
}