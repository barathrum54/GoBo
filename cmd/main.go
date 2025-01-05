package main

import (
	"gobo/internal/app"
	"gobo/internal/cache"
	"gobo/internal/db"
	"gobo/internal/logger"
	"gobo/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Setup initializes the application's dependencies
func Setup() error {
	// Load .env file only if it exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			return err
		}
		log.Println(".env file loaded successfully.")
	} else {
		log.Println("No .env file found, using existing environment variables.")
	}

	// Initialize GORM
	db.ConnectGORM()
	log.Println("Database connection established with GORM.")

	// Run migrations
	err := models.AutoMigrateExamples(db.GormDB)
	if err != nil {
		return err
	}
	log.Println("Database migrations completed.")

	// Initialize Redis
	cache.Connect()
	log.Println("Redis connected.")

	// Initialize logger
	config := logger.DefaultConfig()
	config.Environment = "development"
	config.OutputPaths = []string{"stdout"}
	logger.InitLogger(config)

	logger.Log.Info("Setup completed successfully.")
	return nil
}
func main() {
	if err := Setup(); err != nil {
		log.Fatalf("Application setup failed: %v", err)
	}

	// Start the Fiber app
	application := app.NewApp()

	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
