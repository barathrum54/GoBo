package main_test

import (
	"net/http"
	"testing"
	"time"

	"gobo/internal/app"

	"github.com/stretchr/testify/assert"
)

func TestHTTPServer(t *testing.T) {
	// Start the server in a goroutine
	go func() {
		application := app.NewApp()
		application.Listen(":3000")
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Perform an HTTP request
	resp, err := http.Get("http://localhost:3000")
	assert.NoError(t, err, "Expected no error making GET request to root path")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP status 200")
}
