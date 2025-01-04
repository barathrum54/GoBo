package app

import (
	"gobo/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()

	// Route'ları kaydet
	routes.Register(app)

	return app
}
