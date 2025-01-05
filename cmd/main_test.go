package main

import (
	"context"
	"os"
	"testing"

	"gobo/internal/cache"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
  os.Setenv("REDIS_URL", "localhost:6379")
  os.Setenv("DATABASE_URL", "host=localhost user=admin password=password dbname=gobo port=5432 sslmode=disable")


	// Run the setup function
	err := Setup()
	assert.NoError(t, err, "Setup should complete without errors")

	// Verify Redis connection
	_, err = cache.RedisClient.Ping(context.Background()).Result()
	assert.NoError(t, err, "Redis should be connected")
}
