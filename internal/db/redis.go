package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"rent/internal/config"
)

var RedisDb *redis.Client

func InitRedis() *redis.Client {
	dsn := config.Cfg.Redis
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     dsn.Addr,
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return RedisDb
}
