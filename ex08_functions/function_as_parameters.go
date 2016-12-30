package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	callback(2, Add)

	// example  case using function as parameter is at pkg
	// strings.IndexFunc that return index position
	str := "String 世界"
	f := func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	}
	fmt.Println(strings.IndexFunc(str, f))
}

func Add(a, b int) {
	fmt.Printf("the sum of %d and %d is : %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

