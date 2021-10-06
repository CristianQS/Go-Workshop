package pkg

type ConfigMap struct {
	ApiVersion string         `yaml:"apiVersion"`
	Kind     string   `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Data     Data     `yaml:"data"`
}
