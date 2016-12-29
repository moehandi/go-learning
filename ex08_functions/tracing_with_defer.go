package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
)

func main() {
	MyFunctionTrace()

	MyFunctionTraceLogValues("IKONAH")
}

func MyFunctionTrace() {
	trace("MyFunctionTrace")
	defer untrace("MyFunctionTrace")
	fmt.Println("in MyFunctionTrace")
}
func trace(s string){
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

// function for another example
func MyFunctionTraceLogValues(s string) (n int, err error) {
	defer func() {
		log.Printf("MyFunctionTraceLogValues(%q) = %d , %v", s, n, err)
	}()
	return 7, io.EOF
}