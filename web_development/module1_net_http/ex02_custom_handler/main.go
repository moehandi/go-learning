package main

import (
		"net/http"
		"fmt"
)

type myHandler struct {
		myResponse string
}

func main()  {
		mux := http.NewServeMux()

		// pass myHandler directly to Handle
		mux.Handle("/", &myHandler{"My Message"})

		// OR
		h1 := &myHandler{"This is handler 1"}
		mux.Handle("/h1", h1)

		http.ListenAndServe(":8080",  mux)
}

// any type of object that provide method ServeHTTP from http.Handler like below
// means it implement http.Handler and can become a Handler
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, m.myResponse) // the response will return
}
