package yaml

import (
	"gopkg.in/yaml.v2"
	"log"
)

type YamlV2Serializer struct{}

func (y *YamlV2Serializer) Serialize(object T) (result []byte) {
	result, err := yaml.Marshal(object)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	return result
}

func (y *YamlV2Serializer) Deserialize(body []byte, out interface{})  {
	err := yaml.Unmarshal(body, out)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
}
