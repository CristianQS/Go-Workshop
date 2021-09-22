package yaml

type T interface {}

type YamlSerializer interface {
	Serialize(T) []byte
	Deserialize(body []byte,out interface{})
}
