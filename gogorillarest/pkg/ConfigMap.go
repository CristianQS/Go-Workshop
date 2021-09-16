package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigMap struct {
	ApiVersion string         `yaml:"apiVersion"`
	Kind     string   `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Data     Data     `yaml:"data"`
}

func (c *ConfigMap) GetYaml() *ConfigMap {
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
