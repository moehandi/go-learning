package main

import "fmt"

// prefer declare interface with suffix "-er" or able if word contains er in the end
// It can't have common variable inside, but can have variable type interface
type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

// with this, now Square implement Shape interface
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	square := new(Square)
	square.side = 8
	fmt.Printf("The square Area() is %f\n", square.Area())

	// we can assign square which implement interface to area2
	// now area2 point to square too, and implement interface
	area2 := square
	fmt.Printf("The area2 Area() is %f\n", area2.Area())
}
