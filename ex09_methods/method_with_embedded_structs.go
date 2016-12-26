package main

import "fmt"

type Rectangle struct {
	X float32
	Y float32
}

type Box struct {
	Rectangle // embbding Rectangle type, assume it's edge of box
	Color string
}

func main()  {
	box := Box{Rectangle{8, 8}, "red"}
	fmt.Println("Edge Size", box.Size()) // type Rectangle will promote it's Size method to Box type
	fmt.Println("Box Volume", box.Volume(8)) // volume it's box method has
}

// method for Rectangle struct that Box doesn't have
func (r Rectangle) Size() float32  {
	return r.X * r.Y
}

// method that Box does have
func (b Box) Volume(length float32) float32 {
	var volume float32
	volume = b.Rectangle.Size() * b.X
	return volume
}
