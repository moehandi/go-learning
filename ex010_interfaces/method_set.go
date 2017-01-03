package main

import "fmt"

type List []int

func (l List) Len() int  {
	return len(l)
}

func (l *List) Append(val int)  {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func BiggerThanTen(l Lener) bool  {
	return l.Len() > 10
}

func main()  {
	var bareList List // bare value
	bareList = make([]int, 11) // 11 is > 10
	// not valid: list not pointer List *List
	// and not implement Appender Interface
	// CountInto(list, 1, 1)

	if BiggerThanTen(bareList) { // valid : Identical receiver type, is value not pointer
		fmt.Println("bareList: is long enough")
	}

	// pointer value
	pointerList := new(List)
	CountInto(pointerList, 1, 12)
	if BiggerThanTen(pointerList){ // it still valid,
		// *List can be dereferenced for the receiver
		fmt.Println("pointerList: is Long enough too")
	}

}


