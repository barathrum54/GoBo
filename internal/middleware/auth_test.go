package middleware

import (
	"encoding/base64"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestBasicAuthMiddlewareValidCredentials tests the middleware with valid credentials.
func TestBasicAuthMiddlewareValidCredentials(t *testing.T) {
	// Define valid credentials
	username := "admin"
	password := "password"

	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the BasicAuthMiddleware
	app.Get("/protected", BasicAuthMiddleware(username, password), func(c *fiber.Ctx) error {
		return c.SendString("Authorized")
	})

	// Create an Authorization header with valid credentials
	credentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Basic "+credentials)

	// Perform the request
	resp, err := app.Test(req)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// TestBasicAuthMiddlewareInvalidCredentials tests the middleware with invalid credentials.
func TestBasicAuthMiddlewareInvalidCredentials(t *testing.T) {
	// Define valid credentials
	username := "admin"
	password := "password"

	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the BasicAuthMiddleware
	app.Get("/protected", BasicAuthMiddleware(username, password), func(c *fiber.Ctx) error {
		return c.SendString("Authorized")
	})

	// Create an Authorization header with invalid credentials
	credentials := base64.StdEncoding.EncodeToString([]byte("wronguser:wrongpassword"))
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Basic "+credentials)

	// Perform the request
	resp, err := app.Test(req)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

// TestBasicAuthMiddlewareMissingHeader tests the middleware with no Authorization header.
func TestBasicAuthMiddlewareMissingHeader(t *testing.T) {
	// Define valid credentials
	username := "admin"
	password := "password"

	// Create a new Fiber app
	app := fiber.New()

	// Register a test route with the BasicAuthMiddleware
	app.Get("/protected", BasicAuthMiddleware(username, password), func(c *fiber.Ctx) error {
		return c.SendString("Authorized")
	})

	// Create a request without an Authorization header
	req := httptest.NewRequest("GET", "/protected", nil)

	// Perform the request
	resp, err := app.Test(req)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}
