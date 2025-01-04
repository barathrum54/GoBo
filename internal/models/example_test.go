package models

import (
	"context"
	"gobo/internal/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupTestDB initializes the test database, dropping and recreating the required table
func setupTestDB() {
	log.Println("[Setup] Starting test database setup...")

	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("[Setup] Error loading .env file: %v", err)
	}

	// Connect to the database
	db.Connect()
	log.Println("[Setup] Connected to the database.")

	// Drop the table if it exists
	_, err = db.Conn.Exec(context.Background(), "DROP TABLE IF EXISTS examples")
	if err != nil {
		log.Fatalf("[Setup] Failed to drop table: %v", err)
	}
	log.Println("[Setup] Table 'examples' dropped successfully.")

	// Create the table
	_, err = db.Conn.Exec(context.Background(), `
		CREATE TABLE examples (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("[Setup] Failed to create table: %v", err)
	}
	log.Println("[Setup] Table 'examples' created successfully.")

	log.Println("[Setup] Test database setup completed.")
}

// teardownTestDB cleans up the test database by dropping the test table
func teardownTestDB() {
	log.Println("[Teardown] Starting test database teardown...")

	// Drop the table
	_, err := db.Conn.Exec(context.Background(), "DROP TABLE IF EXISTS examples")
	if err != nil {
		log.Printf("[Teardown] Failed to drop table: %v", err)
	} else {
		log.Println("[Teardown] Table 'examples' dropped successfully.")
	}

	log.Println("[Teardown] Test database teardown completed.")
}

// TestCreateExample tests the CreateExample function for inserting a new record
func TestCreateExample(t *testing.T) {
	log.Println("[TestCreateExample] Starting test...")

	// Setup and teardown
	setupTestDB()
	defer teardownTestDB()

	// Act: Insert a new record
	err := CreateExample("Test Name")
	assert.NoError(t, err, "[TestCreateExample] Error occurred while creating an example.")

	// Assert: Verify the record was inserted
	rows, err := db.Conn.Query(context.Background(), "SELECT id, name FROM examples")
	assert.NoError(t, err, "[TestCreateExample] Error occurred while querying the database.")

	var count int
	for rows.Next() {
		count++
	}

	assert.Equal(t, 1, count, "[TestCreateExample] Expected 1 row in examples table.")
	log.Println("[TestCreateExample] Test completed successfully.")
}

// TestGetExamples tests the GetExamples function for retrieving all records
func TestGetExamples(t *testing.T) {
	log.Println("[TestGetExamples] Starting test...")

	// Setup and teardown
	setupTestDB()
	defer teardownTestDB()

	// Act: Insert multiple records
	CreateExample("Example 1")
	CreateExample("Example 2")

	// Act: Retrieve the records
	examples, err := GetExamples()
	assert.NoError(t, err, "[TestGetExamples] Error occurred while retrieving examples.")

	// Assert: Verify the number of records
	assert.Equal(t, 2, len(examples), "[TestGetExamples] Expected 2 examples.")
	assert.Equal(t, "Example 1", examples[0].Name, "[TestGetExamples] First example name mismatch.")
	assert.Equal(t, "Example 2", examples[1].Name, "[TestGetExamples] Second example name mismatch.")

	log.Println("[TestGetExamples] Test completed successfully.")
}
