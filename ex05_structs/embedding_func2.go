package main

import "fmt"

type Logs struct {
	msg string
}

type Bikes struct {
	Name string
	Logs // embedding Logs struct not as a pointer like ex embedding_func1.go
}

func main()  {
	b := &Bikes{"Mountain Bike", Logs{"1. Mountain bike is awesome"}}
	b.Add("2. Through the hill")
	fmt.Println(b)
}

func (l *Logs) Add(s string)  {
	l.msg += "\n" + s
}

// now bikes have it's own String() method
func (b *Bikes) String() string {
	return b.Name + "\n it's Log" + fmt.Sprintln(b.Logs)
}

func (l *Logs) String() string {
	return l.msg
}