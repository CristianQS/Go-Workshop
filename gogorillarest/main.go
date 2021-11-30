package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gogorillarest/pkg"
	"gogorillarest/pkg/Dtos"
	"gogorillarest/pkg/configmap"
	"gogorillarest/pkg/serializers/yaml"
	"log"
	"net/http"
	"os"
)

var (
	redisConn  = pkg.NewRedisConnection(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"))
	repository = pkg.NewRedisRepository(redisConn.GetRedisConnection())
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/configurations", AddConfigMap()).Methods(http.MethodPost)
	r.HandleFunc("/configurations/{id}", GetConfigMapById()).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func AddConfigMap() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var configurationDto Dtos.ConfigurationDto
		_ = json.NewDecoder(r.Body).Decode(&configurationDto)
		service := configmap.NewService(&yaml.V2Serializer{}, *repository)
		service.AddConfigMap(configurationDto.Id, configurationDto.Value)
		w.WriteHeader(http.StatusCreated)
	}
}

func GetConfigMapById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		service := configmap.NewService(&yaml.V2Serializer{}, *repository)
		configMap := service.GetConfigMapById(vars["id"])
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(configMap)
	}
}

