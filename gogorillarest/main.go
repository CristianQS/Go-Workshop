package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gogorillarest/pkg"
	"gogorillarest/pkg/serializers/yaml"
	"io/ioutil"
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
		bytes, err := ioutil.ReadFile("testdata/configmap.yaml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		serializer := yaml.YamlV2Serializer{}
		configmap := &pkg.ConfigMap{}
		serializer.Deserialize(bytes, configmap)
		repository.Set(vars["id"], *configmap)
		w.WriteHeader(http.StatusCreated)
	}).Methods(http.MethodPost)
	r.HandleFunc("/configuration/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		configMap := repository.GetById(vars["id"])
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(configMap)
	}).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}


