package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected")
}