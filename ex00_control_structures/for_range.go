package main

import "fmt"

func main() {
	str := "Go Code"
	for k, v := range str { // %d will format base-10 formating / decimal
		fmt.Printf("position %d, is %d \n", k, v)
	}

	fmt.Println("index \tint(rune) \trune \t\tchar \tbytes")
	// all v is rune type
	for k, v := range str { // %c wll format to character
		fmt.Printf("%-2d \t%d \t\t%U \t\t%c \t%X \n", k, v, v, v, v)
		//fmt.Println(k, v)
	}

}
