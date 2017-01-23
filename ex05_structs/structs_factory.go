package main

import "fmt"

type Teacher struct {
	Name string
	Age int
}
// it's just for base knowledge see more concrete example
// in example_more/force_use_factory_struct
func main()  {
	t := NewTeacher("teacher", 45)
	fmt.Println(t.Name, t.Age)
}

// then the factory, which returns a pointer to Teacher struct
// if Teacher is defined as a struct type,
// then expression new(Teacher) and &Teacher are equivalent
func NewTeacher(n string, a int) *Teacher  {
	if a < 0 {
		return nil
	}
	return &Teacher{n, a}
}
