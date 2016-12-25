package main

import "fmt"

// Structs 1
type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

// Structs 2
// with this we do embedding, we change with different names
type Point2 struct {
	X, Y int
}

type Circle2 struct {
	Center2 Point2
	Radius int
}

type Wheel2 struct {
	Circles2 Circle2
	Spokes int
}

// STRUCTS 3 same as before but for simplicity we give no name to embedded fields
type Point3 struct {
	X, Y int
}

type Circle3 struct {
	Point3 // no name field, direct to structs Point3
	Radius int
}

type Wheel3 struct {
	Circle3
	Spokes int
}

func main()  {
	// we use structs in STRUCTS 1 section
	var wheel1 Wheel
	wheel1.X = 9
	wheel1.Y = 9
	wheel1.Radius = 5
	wheel1.Spokes = 8
	fmt.Println("wheel structs 1", wheel1)

	// embedding,  we use structs in STRUCTS 2 section
	var wheel2 Wheel2
	wheel2.Circles2.Center2.X = 9 // still long to access within
	wheel2.Circles2.Center2.Y = 9
	wheel2.Circles2.Radius = 5
	wheel2.Spokes = 8
	fmt.Println("wheel structs 2", wheel2)

	// simple than before
	var wheel3 Wheel3
	wheel3.X = 9
	wheel3.Y = 9
	wheel3.Radius = 5
	wheel3.Spokes = 8
	fmt.Println("wheel structs 3", wheel3)
	// with  # adverbs cause printf's %v verb display values in a form to Go syntax
	fmt.Printf("%#v\n",	wheel3)
}
