package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// BasicAuthMiddleware provides basic authentication for routes.
func BasicAuthMiddleware(username, password string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Decode the Base64 encoded credentials
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header",
			})
		}

		// Split the decoded credentials into username and password
		credentials := strings.SplitN(string(decodedCredentials), ":", 2)
		if len(credentials) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		// Validate the credentials
		if credentials[0] != username || credentials[1] != password {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Allow the request to proceed
		return c.Next()
	}
}
