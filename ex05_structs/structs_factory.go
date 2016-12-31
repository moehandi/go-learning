package main

import "fmt"

type Teacher struct {
	Name string
	Age int
}

func main()  {
	t := NewTeacher("teacher", 45)
	fmt.Println(t.Name, t.Age)
}

// then the factory, which returns a pointer to Teacher struct
// if Teache is defined as a struct type,
// then expression new(Teacher) and &Teacher are equivalent
func NewTeacher(n string, a int) *Teacher  {
	if a < 0 {
		return nil
	}
	return &Teacher{n, a}
}
