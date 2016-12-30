package main

import "fmt"

func main()  {
	var v = 28
	fmt.Printf("Integer address of %d is %p \n", v, &v)
	fmt.Println("v == *(&v)", v == *(&v))

	var v2 *int
	v2 = &v
	fmt.Printf("The value at memory location %p is %d\n", v2, *v2)
	*v2 = 299
	fmt.Printf("The value at memory location %p is %d\n", v2, *v2)
}
