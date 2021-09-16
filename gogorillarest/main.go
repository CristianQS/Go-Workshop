package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gogorillarest/pkg"
	"log"
	"net/http"
)

var (
	redisConn = pkg.NewRedisConnection("local-redis",":6379")
	repository = pkg.NewRedisRepository(redisConn.GetRedisConnection())
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/configuration/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		repository.Set(vars["id"], "hello")
	}).Methods(http.MethodPost)
	r.HandleFunc("/configuration/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var configMap pkg.ConfigMap
		configMap.GetYaml()
		fmt.Println(configMap)
		key := repository.GetById(vars["id"])
		fmt.Println(key)
	}).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}


