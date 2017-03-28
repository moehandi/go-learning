package main

import (
	"fmt"
	"time"
)

// create Employee struct
// see there is nothing comma between fields
type Employee struct {
	Id        int
	Name      string
	Address   string
	Birthday  time.Time
	Position  string
	Salary    int
	ManagerId int
}

type Point struct {
	X, Y int
}

// A structs is an aggregate data type that group together zero or more named values
// of arbitrary types as a single entity
func main() {

	// create instance of struct employee named moehandi
	var moehandi Employee

	// lets access it's field, ex:name
	fmt.Println(moehandi.Name) // return empty because not yet initialize or set

	// try to set instance fields
	moehandi.Id = 168
	moehandi.Name = "Moehandi"
	moehandi.Address = "jakarta"
	moehandi.Position = "Manager"
	// and so on another field set with initial value
	fmt.Println(moehandi)

	// struct also can be created with new keyword or &
	newEmployee := new(Employee)
	newEmployee.Id = 123
	newEmployee.Name = "mr.gopher"
	fmt.Println(newEmployee)

	// its called literal of struct, we provide all fields
	newEmployee2 := &Employee{189,"newer","jakarta",time.Now(),"coder", 2000000, 9099000}
	fmt.Println(newEmployee2)


	// kinds of create instance struct
	var point0 Point // same as before
	point0.X = 1
	point0.Y = 2

	var point1 = Point{12, 18}
	var point2 = Point{18, 26}
	point3 := Point{26, 88}

	fmt.Println("point0.x", point0.X, "point0.y", point0.Y)
	fmt.Println("point1.x", point1.X, "point1.y", point1.Y)
	fmt.Println("point2.x", point2.X, "point2.y", point2.Y)
	fmt.Println("point2.x", point3.X, "point3.y", point3.Y)

	// comparing struct
	// if all structs fields is comparable, the structs too
	fmt.Println("point0.x == point1.x", point0.X == point1.X) // false 1 != 12
	fmt.Println("point1.y == point2.x", point1.Y == point2.X) // true 18 == 18

	p1 := &Point{1, 2}
	fmt.Println("p1", p1)
	// equivalent to this, but &Point{1,2} can be used directly within function, such as a function call
	p2 := new(Point)
	*p2 = Point{1,2}
	fmt.Println("p2", p2)

}
