// Package testhelpers contains helper functions for setting up and tearing down the test database.
package testhelpers

import (
	"gobo/internal/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// SetupGormTestDB initializes the test database using GORM.
// It loads the environment variables, connects to the database, and runs migrations.
//
// Parameters:
// - t (*testing.T): The test context for managing test state.
// - models ...interface{}: Variadic parameter for models to migrate.
func SetupGormTestDB(t *testing.T, models ...interface{}) {
	log.Println("[Setup] Starting GORM test database setup...")

	// Load .env variables
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("[Error] Error loading .env file: %v", err)
	}

	// Connect to GORM
	db.ConnectGORM()

	// Run database migrations for each provided model
	for _, model := range models {
		if err := db.GormDB.AutoMigrate(model); err != nil {
			t.Fatalf("[Error] Error during migration for model %v: %v", model, err)
		}
	}

	log.Println("[Setup] Test database setup completed.")
}

// TeardownGormTestDB drops all test tables to clean up after tests.
// It ensures the database is returned to a clean state after testing.
//
// Parameters:
// - models ...interface{}: Variadic parameter for models to drop.
func TeardownGormTestDB(models ...interface{}) {
	log.Println("[Teardown] Starting GORM test database teardown...")

	// Drop each model's table
	for _, model := range models {
		if err := db.GormDB.Migrator().DropTable(model); err != nil {
			log.Printf("[Teardown] Failed to drop table for model %v: %v", model, err)
		} else {
			log.Printf("[Teardown] Table for model %v dropped successfully.", model)
		}
	}

	log.Println("[Teardown] Test database teardown completed.")
}
