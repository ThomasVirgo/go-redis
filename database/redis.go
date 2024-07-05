package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var CTX = context.Background()
var DB *redis.Client

func InitRedisClient() {
	DB = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
