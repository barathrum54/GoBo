// Package routes defines and registers all HTTP routes for the application.
// It integrates with the Fiber web framework and connects to the database layer for handling requests.
package routes

import (
	"gobo/internal/db"
	"gobo/internal/middleware"
	"gobo/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// Response structs for Swagger
type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateExampleResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// Request struct for creating an example
type CreateExampleRequest struct {
	Name string `json:"name"` // The name of the example to be created.
}

// Register registers all routes for the application.
// It maps HTTP endpoints to their corresponding handlers and integrates them with the database.
//
// Parameters:
// - app (*fiber.App): The Fiber application instance to which routes are registered.
func Register(app *fiber.App) {
	// Serve the Swagger documentation at the /swagger endpoint.
	app.Get("/swagger/*", swagger.HandlerDefault) // Default path: /swagger/index.html


	// Root endpoint: Responds with a simple "Hello, World!" message.
	// GET /
	app.Get("/", rootHandler)

	// Retrieve all examples from the database.
	// GET /examples
	app.Get("/examples", getAllExamplesHandler)

	// Group for protected POST routes
	protected := app.Group(
		"/examples",
		middleware.BasicAuthMiddleware("admin", "password"), // Basic Authentication
		middleware.RateLimitMiddleware(10, 1),               // Rate Limiting | x requests per y seconds
	)
	// Create a new example in the database.
	// POST /examples
	protected.Post("/", createExampleHandler)
}

// rootHandler handles the root endpoint.
// @Summary      Root Endpoint
// @Description  Responds with a simple "Hello, World!" message.
// @Tags         root
// @Accept       */*
// @Produce      text/plain
// @Success      200 {string} string "Hello, World!"
// @Router       / [get]
func rootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!") // Respond with a plain text message.
}

// getAllExamplesHandler retrieves all examples from the database and returns them as JSON.
// @Summary      Get All Examples
// @Description  Retrieves all examples from the database.
// @Tags         examples
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Example
// @Failure      500 {object} ErrorResponse
// @Router       /examples [get]
func getAllExamplesHandler(c *fiber.Ctx) error {
	var examples []models.Example // Slice to hold the retrieved examples.

	// Query the database for all examples.
	if result := db.GormDB.Find(&examples); result.Error != nil {
		// Return a 500 status code if there is an error during the query.
		return c.Status(500).JSON(ErrorResponse{Error: "Failed to fetch examples"})
	}

	// Return the examples as a JSON response.
	return c.JSON(examples)
}

// createExampleHandler handles the creation of a new example in the database.
// @Summary      Create Example
// @Description  Creates a new example in the database.
// @Tags         examples
// @Accept       json
// @Produce      json
// @Param        request body      CreateExampleRequest true "Example Request"
// @Success      201 {object} CreateExampleResponse
// @Failure      400 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /examples [post]
func createExampleHandler(c *fiber.Ctx) error {
	var body CreateExampleRequest

	// Parse the JSON request body into the request struct.
	if err := c.BodyParser(&body); err != nil {
		// Return a 400 status code if the request body is invalid.
		return c.Status(400).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	// Create a new example record using the parsed data.
	example := models.Example{Name: body.Name}
	if result := db.GormDB.Create(&example); result.Error != nil {
		// Return a 500 status code if there is an error during record creation.
		return c.Status(500).JSON(ErrorResponse{Error: "Failed to create example"})
	}

	// Convert example.ID from uint to int
	id := int(example.ID)

	// Return a 201 status code and the ID of the newly created example.
	return c.Status(201).JSON(CreateExampleResponse{
		Message: "Example created successfully",
		ID:      id,
	})
}
