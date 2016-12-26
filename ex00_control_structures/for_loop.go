package main

import "fmt"

func main()  {
	forLoop1() // call loop with 1 counter in func
	forLoop2() // call loop with 1+ counter in func
	forLoop3()
	forLoop4()
	forBreak()
	forContinue()
}

// with one counter
func forLoop1() {
	for i := 0; i < 3 ; i++  {
		fmt.Println("here", i)
	}
}

// more than one counter
func forLoop2()  {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}
}

func forLoop3()  {
	char := "*"
	for i := 0; i < 9; i++ {
		fmt.Println(char)
		char += "*"
	}
}

func forLoop4() {
	for i := 0; i < 9; i++ {
		for j := 0; j < i; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func forBreak(){
	for i := 0; i < 10; i++ {
		if i > 5 {
			break // it will break innermost loop then it won't print 6 above
		}
		fmt.Println(i)
	}
}

func forContinue(){
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // it will will skip this part of loop and continue to next iteration
		}
		fmt.Println(i)
	}
}
