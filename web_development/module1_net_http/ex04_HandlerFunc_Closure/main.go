package main

import (
		"net/http"
		"log"
		"fmt"
)


func myFuncHandler (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go function Handler")
}

// func inside func or function closure
func myFuncHandlerClosure(msgResponse string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, msgResponse)
		})
}

func main()  {
		mux := http.NewServeMux()

		mfh := http.HandlerFunc(myFuncHandler)
		mux.Handle("/mfh1", mfh)

		// or with short version
		mux.Handle("/mfh2", http.HandlerFunc(myFuncHandler))

		// see the differences above code with this below
		mux.Handle("/mfhc1", myFuncHandlerClosure("This pass message response"))
		// with another message response
		mux.Handle("/mfhc2", myFuncHandlerClosure("Another Message"))

		fmt.Println("Listening to server on port 8080...")
		// now we can pass http.ListenAndServe as an interface{} that satisfied log.Fatal(... interface{})
		log.Fatal(http.ListenAndServe(":8080", mux))
}


