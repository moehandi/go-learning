// This example show differences between method which passed by pointer and non argument
// see https://golang.org/doc/faq#methods_on_values_or_pointers
package main

import "fmt"

type Boy struct {
	name string
	age  int
}

// Method receiver on value (non-pointer)
func (b Boy) NotChangeName(newName string) {
	b.name = newName
}

// Method receiver on pointer
func (b *Boy) ChangeName(newName string) {
	b.name = newName
}

func main() {
	boy := &Boy{"Rambo", 21}
	fmt.Println("Boy1 Initial Name:", boy.name) // name: Rambo
	boy.ChangeName("Jack")  // it will change name, now name is: Jack
	fmt.Println("Name after ChangeName():", boy.name)
	fmt.Println("Check name again:", boy.name)

	// another example with unchange name
	boy2 := &Boy{"John", 23}
	fmt.Println("Boy2 initial name:", boy2.name)
	boy2.NotChangeName("Brodo")
	fmt.Println("Name after UnchangeName():", boy2.name) // it's name still John, not change to brodo
	fmt.Println("Check name again:", boy2.name)

	fmt.Println("------another example with 'new' keyword on struct------------")
	// now take a look with new keyword
	boy3 := &Boy{"Rambo", 21}
	fmt.Println("Boy1 Initial Name:", boy3.name) // name: Rambo
	boy3.ChangeName("Jack")  // it will change name, now name is: Jack
	fmt.Println("Name after ChangeName():", boy3.name)
	fmt.Println("Check name again:", boy3.name)

	// another example with unchange name
	boy4 := &Boy{"John", 23}
	fmt.Println("Boy2 initial name:", boy4.name)
	boy4.NotChangeName("Brodo")
	fmt.Println("Name after UnchangeName():", boy4.name) // it's name still John, not change to brodo
	fmt.Println("Check name again:", boy4.name)
}
