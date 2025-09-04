package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	// Initialize Redis client here
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Connected to Redis successfully")
}

func GetRedisClient(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

func SetRedisClient(key, value string, expiration time.Duration) error {
    return redisClient.Set(ctx, key, value, expiration).Err()
}

