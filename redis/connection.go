package redis

import "github.com/garyburd/redigo/redis"

func defaultDialFunc() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}

func NewRedis() *Redis {
	return &Redis{
		pool: &redis.Pool{
			MaxIdle:     16,
			MaxActive:   1024,
			IdleTimeout: 300,
			Dial:       defaultDialFunc,
		},
	}
}