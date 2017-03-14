package main

import "fmt"

func main() {
	str := "Go Good"

	str = getKeyValueWithRange(str) // this will return string and fmt.Printf str
	printFormat(str)

}

func getKeyValueWithRange(str string) string {
	for k, v := range str { // %d will format base-10 formating / decimal
		fmt.Printf("pos/idx %d: %d \n", k, v)
	}
	return str
}

func printFormat(str string) {
	fmt.Println("index \tint(rune) \trune \t\tchar \tbytes")
	// all v is rune type
	for k, v := range str { // %c will format to character
		fmt.Printf("%-2d \t%d \t\t%U \t\t%c \t%X \n", k, v, v, v, v)
	}
}
