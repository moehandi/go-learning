package main

import "fmt"

func main() {
	dividedByZero()
	fmt.Println("End of main")
}

func dividedByZero() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Here Recover")
			fmt.Println(err)
			fmt.Println(0)
		}
	}()

	i := 0
	total := 10 / i // will create an error
	fmt.Println("total", total)
}
