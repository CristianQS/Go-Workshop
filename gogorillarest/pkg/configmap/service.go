package configmap

import (
	"gogorillarest/pkg"
	"gogorillarest/pkg/serializers/yaml"
)

type Service interface {
	AddConfigMap(id string, fileData []byte)
	GetConfigMapById(id string) (fileData []byte)
}

type service struct {
	yamlSerializer yaml.YamlSerializer
	repository pkg.ConfigMapRedisRepository
}

func NewService(yamlSerializer yaml.YamlSerializer, repository pkg.ConfigMapRedisRepository) *service {
	return &service{yamlSerializer: yamlSerializer, repository: repository}
}

func (s *service) AddConfigMap(id string, fileData []byte) {
	s.repository.Set(id, fileData)
}

func (s *service) GetConfigMapById(id string) (fileData []byte) {
	return  s.repository.GetById(id)
}
