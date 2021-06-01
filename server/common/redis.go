package common

import (
	e "colte.dev/env"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: e.REDIS_HOST,
	})

	return RedisClient
}
