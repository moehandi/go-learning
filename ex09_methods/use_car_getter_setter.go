package main

import (
	"fmt"
	"github.com/moehandi/go-learning/ex09_methods/car"
)

func main() {
	p := new(car.Car)
	p.SetColor("RED")
	p.SetDoor(4)
	fmt.Println(p.Color(), p.Door()) // now use package car getter and setter to access fields
}
