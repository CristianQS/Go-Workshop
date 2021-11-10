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
)

var (
	redisConn = pkg.NewRedisConnection("local-redis",":6379")
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
		//vars := mux.Vars(r)
		var configurationDto Dtos.ConfigurationDto
		_ = json.NewDecoder(r.Body).Decode(&configurationDto)
		//bytes, err := ioutil.ReadFile(configurationDto.Path)
		//if err != nil {
		//	log.Printf("yamlFile.Get err   #%v ", err)
		//}
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
		w.Write(configMap)
		//_ = json.NewEncoder(w).Encode(configMap)
	}
}



