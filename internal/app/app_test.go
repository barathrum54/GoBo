// Package app_test contains tests for the app package.
// These tests validate the creation and functionality of the Fiber application instance,
// including the dynamic registration of routes.
package app_test

import (
	"testing"

	"gobo/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Test_NewApp verifies the creation of a new Fiber application instance
// and ensures that routes are registered correctly.
// It dynamically collects all registered routes and checks their validity.
func Test_NewApp(t *testing.T) {
	// Create a pseudo Fiber application instance for testing.
	testApp := fiber.New()

	// Register application routes using the routes.Register function.
	routes.Register(testApp)

	// Extract the registered routes from the Fiber application's stack.
	registeredRoutes := testApp.Stack()

	// Dynamically collect the expected routes based on the registered routes.
	expectedRoutes := []string{}
	for methodIndex, routes := range registeredRoutes {
		// Map method indices to their respective HTTP method strings.
		method := map[int]string{
			0: "GET",
			1: "HEAD",
			2: "POST",
			3: "PUT",
			4: "DELETE",
			5: "CONNECT",
			6: "OPTIONS",
			7: "TRACE",
			8: "PATCH",
		}[methodIndex] // Convert index to HTTP method string

		// Append each route's method and path to the expectedRoutes slice.
		for _, route := range routes {
			expectedRoutes = append(expectedRoutes, method+" "+route.Path)
		}
	}

	// Perform assertions for each registered route.
	for _, route := range expectedRoutes {
		// Log the route being checked for debugging purposes.
		t.Logf("Checking route: %s", route)

		// Assert that the expected route is present in the registered routes.
		assert.Contains(t, expectedRoutes, route, "Route not found: "+route)
	}
}
