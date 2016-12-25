package main

import "fmt"

// Both method and function declare with "func keyword"
// the differences is method has method receiver, point to object or type

type Student struct {
	Name string
	Address string
	Age int
}

type Students []Student

func main()  {
	student := Student{"ambo","jakarta", 23}
	ShowInfo(student) // function call
	student.ShowInfo() // method call

	//students := []Student{
	students := Students{
		{"andi","jakarta", 23},
		{"ando","bandung", 25},
	}
	students.ShowInfo()
}

// traditional function
func ShowInfo(s Student){
	fmt.Print("Function call: ")
	fmt.Println(s.Name, s.Address, s.Age)
}

// Declare method with Student type Method receiver
func (s Student) ShowInfo() {
	fmt.Print("Method call: ")
	fmt.Println(s.Name, s.Address, s.Age)
}

// Declare another method with Array, Students type method receiver
func (s Students) ShowInfo(){
	for _, v := range s{
		fmt.Println(v.Name, v.Address, v.Age)
	}
}