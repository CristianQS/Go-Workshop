package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gogorillarest/pkg"
	"gopkg.in/yaml.v2"
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
		repository.Set(vars["id"], "hello")
	}).Methods(http.MethodPost)
	r.HandleFunc("/configuration/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var configMap ConfigMap
		configMap.GetYaml()
		fmt.Println(configMap)
		key := repository.GetById(vars["id"])
		fmt.Println(key)
	}).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (c *ConfigMap) GetYaml() *ConfigMap{
	yamlFile, err := ioutil.ReadFile("testdata/configmap.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile,c)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	return c
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Data struct {
	PlayerInitialLives   string `yaml:"player_initial_lives"`
	UiPropertiesFileName string `yaml:"ui_properties_file_name"`
	GameProperties string `yaml:"game.properties"`
	UserInterfaceProperties string `yaml:"user-interface.properties"`
}

type ConfigMap struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Data Data `yaml:"data"`
}
