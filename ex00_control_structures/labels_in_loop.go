package main

import "fmt"

func main() {
	ContinueLabel()
}

func ContinueLabel() {
HERE:
	for i := 0; i < 6; i++ {
		if i == 3 {
			continue HERE // it will will skip this part of loop and continue to next iteration
		}
		fmt.Println(i)
	}
}
