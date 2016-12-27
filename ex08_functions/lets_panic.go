package main

import "fmt"


func main() {
	// defer and will recover panic, just prints what value panic is
	defer func() {
		fmt.Println( recover() ) // value this_panic
	}()
	panic("this_panic")
}
