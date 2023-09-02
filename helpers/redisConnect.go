package helpers

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisSession() *redis.Client {
	return RedisClient
}

func RedisInit() {
	// Connect to the Redis server
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password for local Redis, set it if needed
		DB:       0,                // Default DB
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Error initializing Redis client: %v", err)
	}
}
