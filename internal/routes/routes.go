package routes

import "github.com/gofiber/fiber/v2"

// Register, uygulama için tüm route'ları kaydeder.
func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Gobo!")
	})

	// Ek route'lar burada tanımlanabilir
}
