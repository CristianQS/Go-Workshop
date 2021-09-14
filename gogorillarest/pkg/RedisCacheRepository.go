package pkg

import "github.com/gomodule/redigo/redis"

type RedisRepository struct{
	conn redis.Conn
}

func NewRedisRepository(conn redis.Conn) *RedisRepository {
	return &RedisRepository{conn: conn}
}

func (r *RedisRepository) Set(key, value string) {
	r.conn.Do("SET", key, value)
}
