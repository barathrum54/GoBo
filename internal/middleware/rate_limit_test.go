package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestRateLimitMiddlewareWithinLimit tests the rate limiter when requests are within the limit.
func TestRateLimitMiddlewareWithinLimit(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the RateLimitMiddleware
	app.Get("/rate-limited", RateLimitMiddleware(10, 1*time.Minute), func(c *fiber.Ctx) error {
		return c.SendString("Request allowed")
	})

	// Perform requests within the limit
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/rate-limited", nil)
		resp, err := app.Test(req)

		// Assert the response
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	}
}

// TestRateLimitMiddlewareExceedLimit tests the rate limiter when requests exceed the limit.
func TestRateLimitMiddlewareExceedLimit(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the RateLimitMiddleware
	app.Get("/rate-limited", RateLimitMiddleware(10, 1*time.Minute), func(c *fiber.Ctx) error {
		return c.SendString("Request allowed")
	})

	// Perform requests to exceed the limit
	for i := 0; i < 12; i++ {
		req := httptest.NewRequest("GET", "/rate-limited", nil)
		resp, err := app.Test(req)

		if i < 10 {
			// Assert the response is allowed within the limit
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)
		} else {
			// Assert the response is blocked after exceeding the limit
			assert.NoError(t, err)
			assert.Equal(t, 429, resp.StatusCode)
		}
	}
}

// TestRateLimitMiddlewareReset tests the rate limiter after the expiration time.
func TestRateLimitMiddlewareReset(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the RateLimitMiddleware
	app.Get("/rate-limited", RateLimitMiddleware(10, 1*time.Second), func(c *fiber.Ctx) error {
		return c.SendString("Request allowed")
	})

	// Perform requests within the limit
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/rate-limited", nil)
		resp, err := app.Test(req)

		// Assert the response is allowed within the limit
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	}

	// Wait for the expiration time to reset the rate limiter
	time.Sleep(2 * time.Second)

	// Perform another request after the reset
	req := httptest.NewRequest("GET", "/rate-limited", nil)
	resp, err := app.Test(req)

	// Assert the response is allowed after reset
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
