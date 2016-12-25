package main

import "fmt"

func main()  {
	var a = 23
	fmt.Println(a, &a)
	p := &a
	*p = 28
	fmt.Println(a, p)
}
