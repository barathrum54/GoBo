// Package routes defines and registers all HTTP routes for the application.
// It integrates with the Fiber web framework and connects to the database layer for handling requests.
package routes

import (
	"gobo/internal/db"
	"gobo/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Register registers all routes for the application.
// It maps HTTP endpoints to their corresponding handlers and integrates them with the database.
//
// Parameters:
// - app (*fiber.App): The Fiber application instance to which routes are registered.
func Register(app *fiber.App) {
	// Root endpoint: Responds with a simple "Hello, World!" message.
	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // Respond with a plain text message.
	})

	// Retrieve all examples from the database.
	// GET /examples
	app.Get("/examples", func(c *fiber.Ctx) error {
		var examples []models.Example // Slice to hold the retrieved examples.

		// Query the database for all examples.
		if result := db.GormDB.Find(&examples); result.Error != nil {
			// Return a 500 status code if there is an error during the query.
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch examples"})
		}

		// Return the examples as a JSON response.
		return c.JSON(examples)
	})

	// Create a new example in the database.
	// POST /examples
	app.Post("/examples", func(c *fiber.Ctx) error {
		// Define the structure for parsing the request body.
		type request struct {
			Name string `json:"name"` // The name of the example to be created.
		}

		var body request

		// Parse the JSON request body into the request struct.
		if err := c.BodyParser(&body); err != nil {
			// Return a 400 status code if the request body is invalid.
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Create a new example record using the parsed data.
		example := models.Example{Name: body.Name}
		if result := db.GormDB.Create(&example); result.Error != nil {
			// Return a 500 status code if there is an error during record creation.
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create example"})
		}

		// Return a 201 status code and the ID of the newly created example.
		return c.Status(201).JSON(fiber.Map{"message": "Example created successfully", "id": example.ID})
	})
}
