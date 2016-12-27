package main

import "fmt"

func main()  {
	// this func not called until reach panic, and then recover is know what to do
	defer func() {// read documentation of recover interface at go doc builtin.recover
		if err := recover(); err != nil {
			fmt.Println("the value from panic is", err) // err = PANIC_404
		}
	}()

	// now panic and send it's value to deferred func recover
	// read documentation of panic interface at go doc builtin.panic
	panic("PANIC_404") // read code above recover

	// this won't happen since panic is
	fmt.Println("check is it show?")
}
