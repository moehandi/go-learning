package main

import (
	"fmt"
)

func main()  {
	classifier(19, -18.933, "Indonesia", complex(1, 8), nil, false)
}

func classifier(args ...interface{})  {
	for i, x := range args {
		switch x.(type) {
		case bool : fmt.Printf("param #%d is a boolean\n", i)
 		case string : fmt.Printf("param #%d is a string\n", i)
		case int, int64 : fmt.Printf("param #%d is a int\n", i)
		case float64 : fmt.Printf("param #%d is a float64\n", i)
		case nil: fmt.Printf("param #%d is a nil\n", i)
		default: fmt.Printf("param #%d unknown type\n", i)
		}
	}
}
