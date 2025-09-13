package connections

import (
	"context"
	"encoding/json"
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

func SaveOTPToRedis(email, otp string) error {
	return SetRedisClient(email, otp, 5*time.Minute)
}

func VarifyOTPFromRedis(email, otp string) (bool, error) {
	storedOTP, err := GetRedisClient(email)
	if err != nil {
		if err == redis.Nil {
			return false, nil // OTP not found
		}
		return false, err // Some other error
	}
	return storedOTP == otp, nil
}

func SaveTempUser(email string, data map[string]string) error {
	b, _ := json.Marshal(data)
	return redisClient.Set(ctx, "tempuser:"+email, b, 5*time.Minute).Err()
}

func GetTempUser(email string) (map[string]string, error) {
	val, err := redisClient.Get(ctx, "tempuser:"+email).Result()
	if err != nil {
		return nil, err
	}
	var user map[string]string
	json.Unmarshal([]byte(val), &user)
	return user, nil
}

func DeleteTempUser(email string) error {
	return redisClient.Del(ctx, "tempuser:"+email).Err()
}

func SaveCartToRedis(key string, cartData interface{}) error {
	b, err := json.Marshal(cartData)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, b, 24*time.Hour).Err()
}

func GetCartFromRedis(key string, cartData *map[string]string) error {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), cartData)
}
