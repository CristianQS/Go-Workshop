package Dtos

type ConfigurationDto struct {
    Id          string  `json:"id"`
	Path        string  `json:"path"`
	Value       []byte `json:"value"`
}