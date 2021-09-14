package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gogorillarest/pkg"
	"log"
	"net/http"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/helloworld", HelloWorldHandler())
	r.HandleFunc("/helloworld/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello %s, you've requested: %s\n", vars["name"], r.URL.Path)
		fmt.Printf("It was requested: %s\n\n", r.URL.Path)
	})
	r.HandleFunc("/configuration/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		redisConn := pkg.NewRedisConnection("local-redis",":6379")
		repository := pkg.NewRedisRepository(redisConn.GetRedisConnection())
		repository.Set(vars["id"], "hello")
	}).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HelloWorldHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Printf("It was requested: %s\n\n", r.URL.Path)
	}
}




