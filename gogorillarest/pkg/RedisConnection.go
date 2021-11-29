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
}

func NewRedisConnection(hostname string, port string) *RedisConnection {
	return &RedisConnection{hostname: hostname, port: port}
}

func (r *RedisConnection) newPool() *redis.Pool {
	addr := *flag.String("addr", r.hostname+":"+r.port, "")
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
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