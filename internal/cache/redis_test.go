// Package cache_test contains tests for the Redis cache utility.
// These tests validate the correctness of Redis connection, and CRUD operations (Set, Get, Delete).
package cache_test

import (
	"os"
	"testing"
	"time"

	"gobo/internal/cache"

	"github.com/stretchr/testify/assert"
)

// TestRedis validates the Redis cache utility functions:
// - Connecting to the Redis server
// - Performing Set, Get, and Delete operations
func TestRedis(t *testing.T) {
	// Set the Redis URL for the test environment.
	// This ensures the test uses the correct Redis server address.
	os.Setenv("REDIS_URL", "localhost:6379")

	// Initialize the Redis connection using the Connect function.
	cache.Connect()

	// Test the Set operation: Add a key-value pair with a 10-second expiration.
	err := cache.Set("test_key", "test_value", 10*time.Second)
	assert.NoError(t, err, "Expected no error during Set operation")

	// Test the Get operation: Retrieve the value of the previously set key.
	value, err := cache.Get("test_key")
	assert.NoError(t, err, "Expected no error during Get operation")
	assert.Equal(t, "test_value", value, "Expected value to match the one set")

	// Test the Delete operation: Remove the key from Redis.
	err = cache.Delete("test_key")
	assert.NoError(t, err, "Expected no error during Delete operation")

	// Verify the key is deleted: Attempt to retrieve the deleted key.
	value, err = cache.Get("test_key")
	assert.Error(t, err, "Expected error when getting a deleted key")
	assert.Empty(t, value, "Expected value to be empty for a deleted key")
}
