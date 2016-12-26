package main

import "fmt"

func main()  {

	for i := 0; i <= 6 ; i++  {
		defer fmt.Println(i)
	}
}
