package main

import (
	"fmt"
)

type Language struct {
	Name string
	Amazing bool
}

func main()  {
	language := Language{"Java", true}
	fmt.Println("before", language.Name)
	language.changeName1("Golang") // it will not change language.Name permanently because pass by value
	fmt.Println("after", language.Name) // still print Java


	language2 := Language{"Java", true}
	fmt.Println("before", language2.Name)
	language2.changeName2("Golang")// it will change language.Name because pass by pointer
	fmt.Println("after", language2.Name) // now print changed to Golang
}

func (l Language) changeName1(newName string) {
	l.Name = newName
}

// method with pointer method receiver (pass by pointer)
// it not pass by value
func (l *Language) changeName2(newName string) {
	l.Name = newName
}



