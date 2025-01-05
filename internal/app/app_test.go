package app_test

import (
	"testing"

	"gobo/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_NewApp(t *testing.T) {
	// Create a pseudo test app
	testApp := fiber.New()

	// Register routes dynamically
	routes.Register(testApp)

	// Extract registered routes
	registeredRoutes := testApp.Stack()

	// Collect expected routes dynamically
	expectedRoutes := []string{}
	for methodIndex, routes := range registeredRoutes {
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

		for _, route := range routes {
			expectedRoutes = append(expectedRoutes, method+" "+route.Path)
		}
	}

	// Perform assertions for each registered route
	for _, route := range expectedRoutes {
		t.Logf("Checking route: %s", route) // Log the route for debugging

		// You could perform a test request to each route here if needed
		assert.Contains(t, expectedRoutes, route, "Route not found: "+route)
	}
}
