package main

import "fmt"

type Animal interface {
	Eat() string
}

type Dog struct {
	food string
}

func (d *Dog) Eat() string {
	return d.food
}

type Cat struct {
	food string
}

func (c *Cat) Eat() string {
	return c.food
}

func main() {

	var anim Animal
	dog := new(Dog)
	dog.food = "meat"

	anim = dog

	// is Dog the type of anim
	if t, ok := anim.(*Dog); ok {
		fmt.Printf("The type of anim is: %T\n", t)
	}

	if u, ok := anim.(*Cat); ok {
		fmt.Printf("The type of anim is: %T\n", u)
	} else {
		fmt.Println("THe anim does not contain a variable of type Cat")
	}

	// another test with type switch
	fmt.Println("With type switch")
	switch t := anim.(type) {
	case *Dog:
		fmt.Printf("Type Dog  %T with value %v\n", t, t)
	case *Cat:
		fmt.Printf("Type Cat  %T with value %v\n", t, t)
	//case string:
	//	fmt.Printf("Type string with value %v\n", t)
	case nil:
		fmt.Printf("nil value, value %v\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

}
