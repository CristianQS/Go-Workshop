package configmap

import (
	"gogorillarest/pkg"
	"gogorillarest/pkg/serializers/yaml"
)

type Service interface {
	GetConfigMap(id string, fileData []byte) (configmap *pkg.ConfigMap)
}

type service struct {
	yamlSerializer yaml.YamlV2Serializer
	repository pkg.ConfigMapRedisRepository
}

func NewService(yamlSerializer yaml.YamlV2Serializer, repository pkg.ConfigMapRedisRepository) *service {
	return &service{yamlSerializer: yamlSerializer, repository: repository}
}

func (s *service) GetConfigMap(id string, fileData []byte) (configmap *pkg.ConfigMap) {
	s.yamlSerializer.Deserialize(fileData, configmap)
	s.repository.Set(id, *configmap)
	return configmap
}
