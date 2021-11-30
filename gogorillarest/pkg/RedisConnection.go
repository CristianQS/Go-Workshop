package pkg

import (
	"context"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

type RedisConnection struct {
	hostname string
	port string
	password string
}

func NewRedisConnection(hostname string, port string, password string) *RedisConnection {
	return &RedisConnection{hostname: hostname, port: port, password: password}
}

func (r *RedisConnection) newPool() *redis.Pool {
	addr := *flag.String("addr", r.hostname+":"+r.port, "")
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", r.password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func (r *RedisConnection) GetRedisConnection() redis.Conn {
	pool := r.newPool()
	conn, err := pool.GetContext(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return conn
}