package main

import (
	"time"
	"fmt"
)

// we can use type like int, float, string, time etc as method receiver
// but because the type (like time is time.Time) not in the same package we must provide an alias to access it
// we can embed as anonymous field in a struct

type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Charset() string {
	return t.Time.String()[0:3]
}

func main()  {
	m := myTime{time.Now()}
	// calling existing String
	fmt.Println("Full time: ", m.String())
	// calling method
	fmt.Println("First 3 char time now", m.first3Charset())
}