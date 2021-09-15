package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisRepository struct{
	conn redis.Conn
}

func NewRedisRepository(conn redis.Conn) *RedisRepository {
	return &RedisRepository{conn: conn}
}

func (r *RedisRepository) Set(key, value string) {
	fmt.Println(key)
	r.conn.Do("SET", key, value)
}

func (r *RedisRepository) GetById(key string) (result string) {
	keys,_ := redis.Bytes(r.conn.Do("GET", key))
	json.Unmarshal(keys, &result)
	return result
}

