package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	pool *redis.Pool
	redisServer = *flag.String("local-redis", ":6379", "")
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
		flag.Parse()
		pool := newPool(redisServer)
		conn, err := pool.GetContext(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.Do("SET", vars["id"], "hello")

	})
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HelloWorldHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Printf("It was requested: %s\n\n", r.URL.Path)
	}
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}


