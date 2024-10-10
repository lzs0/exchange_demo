package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"news_web/global"
)

func InitRedis() {

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis connection failed..")
	}

	global.RedisDb = RedisClient
}
