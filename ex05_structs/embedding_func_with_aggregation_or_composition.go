package main

import "fmt"

type Log struct {
	msg string
}

type Bike struct {
	Name string
	log *Log
}

func main()  {
	b := new(Bike)
	b.Name = "Mountain Bike"
	b.log = new(Log)
	b.log.msg = "1. Mountain bike is awesome"

	// shorter version
	// b := &Bike{"Mountain Bike", &Log{"1. Through the hill"}}
	b.Log().Add("2. Through the hill")
	fmt.Println(b.Log())
}

func (l *Log) Add(s string)  {
	l.msg += "\n" + s
}

func (l *Log) String() string  {
	return l.msg
}

func (b *Bike) Log() *Log  {
	return b.log
}
