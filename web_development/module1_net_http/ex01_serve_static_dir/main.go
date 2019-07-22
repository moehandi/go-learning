package main

import (
	"fmt"
	"net/http"
)

func main() {
		mux := http.NewServeMux()
		// http.FileServer return Handler
		fs := http.FileServer(http.Dir("public"))
		mux.Handle("/", fs)
		fmt.Println("listening server...")
		http.ListenAndServe(":8080", mux)
}
