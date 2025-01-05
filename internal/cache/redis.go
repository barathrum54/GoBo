// Package cache provides utilities for interacting with a Redis server.
// It includes functions for connecting to Redis and performing CRUD operations.
package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient is the global Redis client used for interacting with the Redis server.
var RedisClient *redis.Client

// ctx is the global context used for Redis operations.
var ctx = context.Background()

// Connect initializes the Redis client and establishes a connection to the Redis server.
// It retrieves the Redis server address from the REDIS_URL environment variable.
// If the connection fails, the application terminates with a fatal log.
func Connect() {
	// Retrieve the Redis server address from environment variables.
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		// Default to localhost if no environment variable is set.
		redisAddr = "localhost:6379"
	}

	// Initialize the Redis client with the specified options.
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Address of the Redis server
		Password: "",        // Password for Redis authentication (if any)
		DB:       0,         // Default database index
	})

	// Test the Redis connection using the PING command.
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		// Log a fatal error and terminate if the connection fails.
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Log a message indicating a successful connection.
	log.Println("Connected to Redis!")
}

// Set stores a key-value pair in Redis with an expiration time.
//
// Parameters:
// - key (string): The key to store.
// - value (string): The value to associate with the key.
// - expiration (time.Duration): The time-to-live for the key-value pair.
//
// Returns:
// - error: An error if the operation fails.
func Set(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get retrieves the value associated with a key from Redis.
//
// Parameters:
// - key (string): The key to retrieve.
//
// Returns:
// - string: The value associated with the key.
// - error: An error if the operation fails or the key does not exist.
func Get(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// Delete removes a key from Redis.
//
// Parameters:
// - key (string): The key to delete.
//
// Returns:
// - error: An error if the operation fails or the key does not exist.
func Delete(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
