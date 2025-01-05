package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimitMiddleware creates a rate limiter middleware for Fiber with dynamic parameters.
func RateLimitMiddleware(maxRequests int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        maxRequests,       // Maximum number of requests per duration
		Expiration: expiration,        // Time duration for the limit
		KeyGenerator: func(c *fiber.Ctx) string {
			// Use client IP as the key
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			// Response when rate limit is exceeded
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Try again later.",
			})
		},
	})
}
