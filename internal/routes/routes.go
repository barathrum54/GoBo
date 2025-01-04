package routes

import (
	"gobo/internal/db"
	"gobo/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Register registers all routes for the application
func Register(app *fiber.App) {
	// GET /examples - Retrieve all examples
	app.Get("/examples", func(c *fiber.Ctx) error {
		var examples []models.Example
		if result := db.GormDB.Find(&examples); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch examples"})
		}
		return c.JSON(examples)
	})

	// POST /examples - Create a new example
	app.Post("/examples", func(c *fiber.Ctx) error {
		type request struct {
			Name string `json:"name"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		example := models.Example{Name: body.Name}
		if result := db.GormDB.Create(&example); result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create example"})
		}

		return c.Status(201).JSON(fiber.Map{"message": "Example created successfully", "id": example.ID})
	})
}
