package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

// Connect initializes the Redis client
func Connect() {
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // Varsayılan Redis adresi
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Redis sunucusunun adresi
		Password: "",        // Parola (varsa)
		DB:       0,         // Varsayılan DB
	})

	// Redis bağlantısını test et
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}

// Set sets a key-value pair in Redis with expiration
func Set(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key from Redis
func Get(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// Delete removes a key from Redis
func Delete(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
