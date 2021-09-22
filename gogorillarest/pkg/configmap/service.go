package configmap

import (
	"gogorillarest/pkg"
	"gogorillarest/pkg/serializers/yaml"
)

type Service interface {
	AddConfigMap(id string, fileData []byte) (configmap *pkg.ConfigMap)
	GetConfigMapById(id string) (configmap pkg.ConfigMap)
}

type service struct {
	yamlSerializer yaml.YamlSerializer
	repository pkg.ConfigMapRedisRepository
}

func NewService(yamlSerializer yaml.YamlSerializer, repository pkg.ConfigMapRedisRepository) *service {
	return &service{yamlSerializer: yamlSerializer, repository: repository}
}

func (s *service) AddConfigMap(id string, fileData []byte) (configmap *pkg.ConfigMap) {
	s.yamlSerializer.Deserialize(fileData, &configmap)
	s.repository.Set(id, *configmap)
	return configmap
}

func (s *service) GetConfigMapById(id string) (configmap pkg.ConfigMap) {
	return  s.repository.GetById(id)
}
