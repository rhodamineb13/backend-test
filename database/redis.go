package database

import (
	"github.com/redis/go-redis/v9"

	"github.com/rhodamineb13/backend-test/utils"
)

var RedisDB *redis.Client

func ConnectRedis() {
	opt := &redis.Options{
		Addr: "localhost:6379",
	}
	if utils.RedisUsername != "" && utils.RedisPassword != "" {
		opt.Username = utils.RedisUsername
		opt.Password = utils.RedisPassword
	}
	rdb := redis.NewClient(opt)
	RedisDB = rdb
}
