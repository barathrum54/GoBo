package routes

import (
	"gobo/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Register registers all routes for the application
func Register(app *fiber.App) {
	app.Get("/examples", func(c *fiber.Ctx) error {
		examples, err := models.GetExamples()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch examples"})
		}
		return c.JSON(examples)
	})

	app.Post("/examples", func(c *fiber.Ctx) error {
		type request struct {
			Name string `json:"name"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if err := models.CreateExample(body.Name); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create example"})
		}

		return c.Status(201).JSON(fiber.Map{"message": "Example created successfully"})
	})
}
