package main

import (
		"net/http"
		"fmt"
		"log"
)

func messageHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GO")
}

func main() {
		// watch we dont use NewServeMux, it's means we use default ServeMux
		// and we dont pass mux to the http.ListenAndServe(":8080", mux_here)
		http.HandleFunc("/", messageHandleFunc)

		log.Println("Listening to http: port 8080")
		http.ListenAndServe(":8080", nil)
}
