// Package db_test contains tests for the database connection and configuration.
// These tests validate the GORM connection setup, environment variable handling, and connection pool settings.
package db_test

import (
	"log"
	"os"
	"testing"

	"gobo/internal/db"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the database connection for testing.
// It loads environment variables from the .env file and connects to the database using GORM.
//
// Parameters:
// - t (*testing.T): The test context for managing test state.
func setupGormTestDB(t *testing.T) {
	log.Println("[Setup] Starting test database setup...")

	// Load .env variables
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("[Error] Error loading .env file: %v", err)
	}

	// Connect to the database
	db.ConnectGORM()
	log.Println("[Setup] Test database setup completed.")
}

// teardownGormTestDB cleans up the database connection after testing.
// It closes the SQL DB connection to release resources.
//
// Parameters:
// - t (*testing.T): The test context for managing test state.
func teardownGormTestDB(t *testing.T) {
	log.Println("[Teardown] Starting test database teardown...")

	// Retrieve the SQL DB instance from GORM
	sqlDB, err := db.GormDB.DB()
	if err != nil {
		t.Fatalf("[Error] Failed to retrieve SQL DB instance: %v", err)
	}

	// Close the SQL DB connection
	if err := sqlDB.Close(); err != nil {
		t.Fatalf("[Error] Failed to close SQL DB connection: %v", err)
	}

	log.Println("[Teardown] Test database teardown completed.")
}

// TestConnectGORM_Success validates the successful connection to the database using GORM.
// It ensures that the GORM DB instance and the SQL DB connection are correctly initialized.
func TestConnectGORM_Success(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB(t)

	// Verify that GormDB is initialized
	assert.NotNil(t, db.GormDB, "GormDB should not be nil after connection")

	// Retrieve the SQL DB instance and verify its initialization
	sqlDB, err := db.GormDB.DB()
	assert.NoError(t, err, "Should be able to retrieve SQL DB instance")
	assert.NotNil(t, sqlDB, "SQL DB instance should not be nil")
}

// TestConnectGORM_MissingEnv tests the behavior of the database connection when the DATABASE_URL environment variable is missing.
// It expects a panic to occur during the connection attempt.
func TestConnectGORM_MissingEnv(t *testing.T) {
	// Unset the DATABASE_URL environment variable
	os.Unsetenv("DATABASE_URL")

	// Defer a recovery function to handle the expected panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("[Recovered] Expected panic for missing DATABASE_URL: %v", r)
		} else {
			t.Errorf("Expected a panic for missing DATABASE_URL, but did not get one")
		}
	}()

	// Attempt to connect to the database (a panic is expected)
	db.ConnectGORM()
}

// TestConnectionPoolSettings validates the database connection pool settings.
// It ensures that the number of idle and open connections adheres to the defined limits.
func TestConnectionPoolSettings(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB(t)

	// Retrieve the SQL DB instance
	sqlDB, err := db.GormDB.DB()
	assert.NoError(t, err, "Should be able to retrieve SQL DB instance")

	// Validate MaxIdleConns
	expectedMaxIdleConns := 10
	actualMaxIdleConns := sqlDB.Stats().Idle // Current number of idle connections
	assert.LessOrEqual(t, actualMaxIdleConns, expectedMaxIdleConns, "Idle connections should not exceed MaxIdleConns")

	// Validate MaxOpenConns
	expectedMaxOpenConns := 100
	assert.LessOrEqual(t, sqlDB.Stats().OpenConnections, expectedMaxOpenConns, "Open connections should not exceed MaxOpenConns")
}
