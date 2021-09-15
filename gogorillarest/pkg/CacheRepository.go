package pkg

type Repository interface {
	Set(key, value string)
	GetById(key string) string
}
