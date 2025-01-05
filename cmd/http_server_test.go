// Package main_test contains tests for the main application functionality.
// It includes integration tests for the HTTP server.
package main_test

import (
	"net/http"
	"testing"
	"time"

	"gobo/internal/app"

	"github.com/stretchr/testify/assert"
)

// TestHTTPServer validates the behavior of the HTTP server.
// It ensures the server starts correctly and responds with the expected status code.
func TestHTTPServer(t *testing.T) {
	// Start the Fiber application in a separate goroutine to allow parallel execution.
	go func() {
		// Initialize the application instance using app.NewApp().
		application := app.NewApp()

		// Start the application on port 3000. No error handling here as it's for testing.
		application.Listen(":3000")
	}()

	// Allow the server some time to start up before sending requests.
	time.Sleep(1 * time.Second)

	// Perform an HTTP GET request to the root endpoint.
	resp, err := http.Get("http://localhost:3000")

	// Assert no error occurred while making the HTTP request.
	assert.NoError(t, err, "Expected no error making GET request to root path")

	// Assert that the HTTP response status code is 200 (OK).
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP status 200")
}
