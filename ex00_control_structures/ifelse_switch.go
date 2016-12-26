package main

import "fmt"

func main()  {
	condition := true

	PrintIfElse(condition)

	PrintIfElse2(88)
	PrintIfElse2(168)

	PrintSwitch("golang")
	PrintSwitch("javascript")

	x := []int{2,3,4,5,6}
	y := []int{1,2,3,4,5}
	PrintSwitch2(x, y)
}

func PrintIfElse(condition bool) {
	if condition {
		fmt.Println("condition is true")
	} else { // don't put else { in new line
		fmt.Println("conditon is false")
	}
}

func PrintIfElse2(value int) {
	max := 100
	// rather than write:
	// v := value
	// if v > max { ... }
	// write like below in oneline
	if v := value; v > max {
		fmt.Println("value greater than max")
	} else {
		fmt.Println("value less than max")
	}
}

func PrintSwitch(word string) {
	switch word {
	// case "golang" :
	case "golang","go": // it can be multiple to
		fmt.Println("Golang")
	case "java":
		fmt.Println("Java")
	case "nodejs":
		fmt.Println("NodeJS")
	default:
		fmt.Println("no case, it's default")
	}
}

// example of switch with brief statement
func PrintSwitch2(x []int, y[]int){
	i, j := 1, 2

	// statement like this is valid
	switch a, b := x[i], y[j]; {
	case a < b :
		fmt.Println("a < b")
	case a == b :
		fmt.Println("a = b")
	case a > b:
		fmt.Println("a > b")
	}
}


