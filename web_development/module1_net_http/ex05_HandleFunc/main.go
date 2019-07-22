package main

import (
	"fmt"
	"log"
	"net/http"
)


func myHandlerFunc (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go function Handler")
}


func main()  {
		mux := http.NewServeMux()

		mfh := http.HandlerFunc(myHandlerFunc)
		mux.Handle("/mfh1", mfh)

		// or with short version
		mux.Handle("/mfh2", http.HandlerFunc(myHandlerFunc))

		// this code like short for code above
		// In the previous section, a normal function was converted into a HandlerFunc type and used as an HTTP
		// handler by registering it with ServeMux.Handle. Because ordinary functions are frequently used as HTTP
		// handlers in this way, the http package provides a shortcut method: ServeMux.HandleFunc. The HandleFunc
		// registers the handler function for the given pattern. (This is just a shortcut method for your convenience.)
		// It internally (inside the http package) converts into a HandlerFunc type and registers the handler into
		// ServeMux.
		mux.HandleFunc("/mfh3", myHandlerFunc)

		fmt.Println("Listening to server on port 8080...")
		// now we can pass http.ListenAndServe as an interface{} that satisfied log.Fatal(... interface{})
		log.Fatal(http.ListenAndServe(":8080", mux))
}


