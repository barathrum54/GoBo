// Package main contains the tests for the application's main setup process.
// These tests ensure the application dependencies (e.g., Redis, database) are configured correctly.
package main

import (
	"context"
	"os"
	"testing"

	"gobo/internal/cache"

	"github.com/stretchr/testify/assert"
)

// TestSetup validates the application's setup process.
// It ensures that environment variables are correctly loaded and all services
// (e.g., Redis) are initialized without errors.
func TestSetup(t *testing.T) {
	// Mock environment variables to simulate a configuration for testing purposes.
	// These variables are required for the setup process to succeed.
	os.Setenv("REDIS_URL", "localhost:6379")
	os.Setenv("DATABASE_URL", "host=localhost user=admin password=password dbname=gobo port=5432 sslmode=disable")

	// Run the Setup function to initialize application dependencies.
	err := Setup()
	assert.NoError(t, err, "Setup should complete without errors")

	// Verify the Redis connection by sending a ping command to the Redis server.
	_, err = cache.RedisClient.Ping(context.Background()).Result()
	assert.NoError(t, err, "Redis should be connected")
}
