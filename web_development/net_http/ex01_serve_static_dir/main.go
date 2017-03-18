package main

import (
		"net/http"
		"fmt"
)

func main() {
		mux := http.NewServeMux()
		// http.FileServer return Handler
		fs := http.FileServer(http.Dir("public"))
		mux.Handle("/", fs)
		fmt.Println("listening server...")
		http.ListenAndServe(":8080", mux)
}
