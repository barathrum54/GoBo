package cache_test

import (
	"os"
	"testing"
	"time"

	"gobo/internal/cache"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	// Set up Redis URL for testing
	os.Setenv("REDIS_URL", "localhost:6379")

	// Connect to Redis
	cache.Connect()

	// Test Set operation
	err := cache.Set("test_key", "test_value", 10*time.Second)
	assert.NoError(t, err, "Expected no error during Set operation")

	// Test Get operation
	value, err := cache.Get("test_key")
	assert.NoError(t, err, "Expected no error during Get operation")
	assert.Equal(t, "test_value", value, "Expected value to match the one set")

	// Test Delete operation
	err = cache.Delete("test_key")
	assert.NoError(t, err, "Expected no error during Delete operation")

	// Verify the key is deleted
	value, err = cache.Get("test_key")
	assert.Error(t, err, "Expected error when getting a deleted key")
	assert.Empty(t, value, "Expected value to be empty for a deleted key")
}
