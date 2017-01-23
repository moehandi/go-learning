package main

import (
	persons "github.com/moehandi/go-learning/example_more/force_use_factory_struct/person"
	"fmt"
)

func main()  {
	// person1 := new(persons.person) NOT Valid because person is private
	person1 := persons.New("Moehandi",27)
	fmt.Println("Person: ", person1.Name, person1.Age)
}
