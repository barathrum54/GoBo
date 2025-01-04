package models

import (
	"gobo/internal/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the test database using GORM
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

// teardownGormTestDB drops all test tables
func teardownGormTestDB() {
	log.Println("[Teardown] Starting GORM test database teardown...")

	if err := db.GormDB.Migrator().DropTable(&Example{}); err != nil {
		log.Printf("[Teardown] Failed to drop table: %v", err)
	} else {
		log.Println("[Teardown] Table 'examples' dropped successfully.")
	}

	log.Println("[Teardown] Test database teardown completed.")
}

func TestCreateExampleGorm(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB()

	// Create a new record
	example := Example{Name: "Test Name"}
	result := db.GormDB.Create(&example)
	assert.NoError(t, result.Error, "Error occurred while creating an example")

	// Check database for the number of records
	var count int64
	db.GormDB.Model(&Example{}).Count(&count)
	assert.Equal(t, int64(1), count, "Expected 1 row in examples table")
}

func TestGetExamplesGorm(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB()

	// Add test data
	db.GormDB.Create(&Example{Name: "Example 1"})
	db.GormDB.Create(&Example{Name: "Example 2"})

	// Retrieve records from the table
	var examples []Example
	result := db.GormDB.Find(&examples)
	assert.NoError(t, result.Error, "Error occurred while retrieving examples")

	// Verify the number of records
	assert.Equal(t, 2, len(examples), "Expected 2 examples")
	assert.Equal(t, "Example 1", examples[0].Name)
	assert.Equal(t, "Example 2", examples[1].Name)
}
