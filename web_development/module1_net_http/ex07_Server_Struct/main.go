package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func messageHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GO")
}

func main() {

		http.HandleFunc("/", messageHandleFunc)

		// example before we don't customize server
		// with http.Serve struct we can do it
		server := &http.Server{
				Addr:	":8080",
				ReadTimeout: 10 * time.Second,
				WriteTimeout: 10 * time.Second,
				MaxHeaderBytes: 1 << 20,
		}
		log.Println("Listening to ", server.Addr)
		server.ListenAndServe() // now we dont pass port and mux like before
}
