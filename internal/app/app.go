// Package app is responsible for initializing the Fiber application instance.
// It integrates routes and provides a central entry point for the HTTP server setup.
package app

import (
	"gobo/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// NewApp initializes and returns a new Fiber application instance.
// This function sets up the application with all registered routes.
//
// Returns:
// - *fiber.App: The initialized Fiber application instance ready to handle requests.
func NewApp() *fiber.App {
	// Create a new instance of Fiber.
	app := fiber.New()

	// Register application routes.
	// The routes are defined and handled in the routes package.
	routes.Register(app)

	// Return the initialized Fiber application instance.
	return app
}
