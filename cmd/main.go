// Package main serves as the entry point for the application.
// It handles the initialization of dependencies and starts the HTTP server.
package main

import (
	_ "gobo/docs"
	"gobo/internal/app"
	"gobo/internal/cache"
	"gobo/internal/db"
	"gobo/internal/logger"
	"gobo/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Setup initializes the application's dependencies, including:
// - Loading environment variables
// - Connecting to the database (GORM)
// - Running database migrations for all models
// - Initializing Redis
// - Setting up the logger
// Returns an error if any step in the initialization fails.
func Setup() error {
	// Load the .env file only if it exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			// Return an error if the .env file exists but fails to load
			return err
		}
		log.Println(".env file loaded successfully.")
	} else {
		// Log a message if the .env file is not found
		log.Println("No .env file found, using existing environment variables.")
	}

	// Initialize the database connection using GORM
	db.ConnectGORM()
	log.Println("Database connection established with GORM.")

	// Run database migrations for all models
	err := AutoMigrateAllModels(db.GormDB)
	if err != nil {
		// Return an error if migrations fail
		return err
	}
	log.Println("Database migrations completed.")

	// Initialize Redis connection
	cache.Connect()
	log.Println("Redis connected.")

	// Initialize the application logger
	config := logger.DefaultConfig()
	config.Environment = "development"        // Set logger environment to development
	config.OutputPaths = []string{"stdout"}   // Log output to standard output
	logger.InitLogger(config)

	// Log a message indicating that setup was successful
	logger.Log.Info("Setup completed successfully.")
	return nil
}

// AutoMigrateAllModels migrates all the models automatically using GORM.
// It accepts the GORM DB connection as a parameter and migrates each model in the list.
func AutoMigrateAllModels(db *gorm.DB) error {
	// List all models you want to migrate
	models := []interface{}{
		&models.Example{},
		&models.User{},  // Add other models here as needed
	}

	// Loop through all models and run AutoMigrate
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

// main is the entry point for the application.
// It performs setup, starts the HTTP server, and handles fatal errors.
func main() {
	// Run the setup process and handle any errors
	if err := Setup(); err != nil {
		log.Fatalf("Application setup failed: %v", err)
	}

	// Initialize and start the Fiber HTTP server
	application := app.NewApp()

	// Log the server startup message and listen for incoming requests
	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		// Log and terminate the application if the server fails to start
		log.Fatalf("Error starting server: %v", err)
	}
}
