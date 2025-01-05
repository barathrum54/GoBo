// Package models contains the application's database models and related functionality.
// This file includes tests for the Example model using GORM.
package models

import (
	"gobo/internal/db"
	"gobo/internal/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateExampleGorm validates the creation of an Example record in the database.
// It ensures that the record is successfully inserted and counted in the table.
func TestCreateExampleGorm(t *testing.T) {
	// Set up the test database with Example model
	testhelpers.SetupGormTestDB(t, &Example{})
	defer testhelpers.TeardownGormTestDB(&Example{})

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
	// Set up the test database with Example model
	testhelpers.SetupGormTestDB(t, &Example{})
	defer testhelpers.TeardownGormTestDB(&Example{})

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
