package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type ConfigMapRedisRepository struct{
	conn redis.Conn
}

func NewRedisRepository(conn redis.Conn) *ConfigMapRedisRepository {
	return &ConfigMapRedisRepository{conn: conn}
}

func (r *ConfigMapRedisRepository) Set(key string, value ConfigMap) {
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	r.conn.Do("SET", key, bytes)
}

func (r *ConfigMapRedisRepository) GetById(key string) (result ConfigMap) {
	keys,_ := redis.Bytes(r.conn.Do("GET", key))
	json.Unmarshal(keys, &result)
	return result
}

