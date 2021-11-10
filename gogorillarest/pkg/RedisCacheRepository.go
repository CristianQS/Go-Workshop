package pkg

import (
	"github.com/gomodule/redigo/redis"
)

type ConfigMapRedisRepository struct{
	conn redis.Conn
}

func NewRedisRepository(conn redis.Conn) *ConfigMapRedisRepository {
	return &ConfigMapRedisRepository{conn: conn}
}

func (r *ConfigMapRedisRepository) Set(key string, value []byte) {
	//bytes, err := json.Marshal(value)
	//if err != nil {
	//	fmt.Println(err)
	//}
	r.conn.Do("SET", key, value)
}

func (r *ConfigMapRedisRepository) GetById(key string) []byte {
	keys,_ := redis.Bytes(r.conn.Do("GET", key))
	return keys
}

