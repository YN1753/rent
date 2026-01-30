package db

import (
	"context"
	"rent/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	dsn := config.Cfg.Redis
	RedisDb := redis.NewClient(&redis.Options{
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
