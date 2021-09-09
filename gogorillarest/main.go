package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/helloworld", HelloWorldHandler())
	r.HandleFunc("/helloworld/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello %s, you've requested: %s\n", vars["name"], r.URL.Path)
		fmt.Printf("It was requested: %s\n\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HelloWorldHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Printf("It was requested: %s\n\n", r.URL.Path)
	}
}