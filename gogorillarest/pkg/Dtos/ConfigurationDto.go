package Dtos

type ConfigurationDto struct {
    Id          string  `json:"id"`
	Value       []byte `json:"value"`
}