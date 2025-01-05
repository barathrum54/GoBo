// Package models contains the application's database models and related functionality.
// This file includes tests for the Example model using GORM.
package models

import (
	"gobo/internal/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the test database using GORM.
// It loads the environment variables, connects to the database, and runs migrations.
//
// Parameters:
// - t (*testing.T): The test context for managing test state.
func setupGormTestDB(t *testing.T) {
	log.Println("[Setup] Starting GORM test database setup...")

	// Load .env variables
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("[Error] Error loading .env file: %v", err)
	}

	// Connect to GORM
	db.ConnectGORM()

	// Run database migrations and check for errors
	if err := AutoMigrateExamples(db.GormDB); err != nil {
		t.Fatalf("[Error] Error during migrations: %v", err)
	}

	log.Println("[Setup] Test database setup completed.")
}

// teardownGormTestDB drops all test tables to clean up after tests.
// It ensures the database is returned to a clean state after testing.
func teardownGormTestDB() {
	log.Println("[Teardown] Starting GORM test database teardown...")

	// Drop the "examples" table
	if err := db.GormDB.Migrator().DropTable(&Example{}); err != nil {
		log.Printf("[Teardown] Failed to drop table: %v", err)
	} else {
		log.Println("[Teardown] Table 'examples' dropped successfully.")
	}

	log.Println("[Teardown] Test database teardown completed.")
}

// TestCreateExampleGorm validates the creation of an Example record in the database.
// It ensures that the record is successfully inserted and counted in the table.
func TestCreateExampleGorm(t *testing.T) {
	// Set up the test database
	setupGormTestDB(t)
	defer teardownGormTestDB()

	// Create a new record
	example := Example{Name: "Test Name"}
	result := db.GormDB.Create(&example)
	assert.NoError(t, result.Error, "Error occurred while creating an example")

	// Check the number of records in the examples table
	var count int64
	db.GormDB.Model(&Example{}).Count(&count)
	assert.Equal(t, int64(1), count, "Expected 1 row in examples table")
}

// TestGetExamplesGorm validates the retrieval of Example records from the database.
// It ensures that records are successfully retrieved and match the expected values.
func TestGetExamplesGorm(t *testing.T) {
	// Set up the test database
	setupGormTestDB(t)
	defer teardownGormTestDB()

	// Add test data to the examples table
	db.GormDB.Create(&Example{Name: "Example 1"})
	db.GormDB.Create(&Example{Name: "Example 2"})

	// Retrieve all records from the examples table
	var examples []Example
	result := db.GormDB.Find(&examples)
	assert.NoError(t, result.Error, "Error occurred while retrieving examples")

	// Verify the number of records and their content
	assert.Equal(t, 2, len(examples), "Expected 2 examples")
	assert.Equal(t, "Example 1", examples[0].Name)
	assert.Equal(t, "Example 2", examples[1].Name)
}
