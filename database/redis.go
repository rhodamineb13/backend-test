package database

import (
	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	RedisDB = rdb
}
