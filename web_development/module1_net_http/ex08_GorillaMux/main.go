package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetApi(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the GO")
}

func main() {

		r := mux.NewRouter().StrictSlash(false)
		r.HandleFunc("/api/getapi", GetApi).Methods("GET")

		// provide another handler for code below
		//r.HandleFunc("/api/notes", postapi).Methods("POST")
		//r.HandleFunc("/api/notes/{id}", putapi).Methods("PUT")
		//r.HandleFunc("/api/notes/{id}", deleteapi).Methods("DELETE")

		server := &http.Server{
				Addr:
				":8080",
				Handler: r, // we assign r (route) mux to the handler
		}

		log.Println("Listening to ", server.Addr)
		server.ListenAndServe() // now we don't pass port and mux like before
}
