package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler  struct {
		response string
}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, m.response)
}

func myFuncHandler (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go function Handler")
}

func main()  {
		mux := http.NewServeMux()

		// this is we provide type implement http.Handler
		mh1 := &myHandler{"Response from mh1"} // in this line we provide response string with our own string text
		mux.Handle("/mh1", mh1)

		// Instead of creating custom handler types by implementing the http.Handler interface, you can use the
		// http.HandlerFunc type to serve as an HTTP handler. You can convert any function into a HandlerFunc type
		// if the function has the signature func(http.ResponseWriter, *http.Request). The HandlerFunc type
		// works as an adapter that allows you to use normal functions as HTTP handlers. The HandlerFunc type has a
		// built-in method ServeHTTP(http.ResponseWriter, *http.Request), so it also satisfies the http.Handler
		// interface and can work as an HTTP handler.

		// we pass func as a handler when func satisfied signature like ServeHTTP
		// but with this case we can't pass Response string object
		// see another example in ex04_HandlerFunc_Closure
		mfh := http.HandlerFunc(myFuncHandler)
		mux.Handle("/mfh1", mfh)

		// or with short version
		mux.Handle("/mfh2", http.HandlerFunc(myFuncHandler))

		fmt.Println("Listening to server on port 8080...")
		// now we can pass http.ListenAndServe as an interface{} that satisfied log.Fatal(... interface{})
		log.Fatal(http.ListenAndServe(":8080", mux))
}
