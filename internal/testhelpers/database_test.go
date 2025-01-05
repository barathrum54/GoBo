// Package testhelpers contains helper functions for setting up and tearing down the test database.
// This file includes tests for the database setup, migration, and teardown.
package testhelpers

import (
	"gobo/internal/db"
	"gobo/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDatabaseConnection ensures that the database connection is established correctly.
func TestDatabaseConnection(t *testing.T) {
	// Set up the test database with models to test migrations
	SetupGormTestDB(t, &models.User{}, &models.Example{})
	defer TeardownGormTestDB(&models.User{}, &models.Example{})

	// Test that GORM is connected and the models are migrated correctly
	assert.NotNil(t, db.GormDB, "Database connection should not be nil")

	// Optionally, check if some records exist after migrations
	var count int64
	db.GormDB.Model(&models.User{}).Count(&count)
	assert.Equal(t, int64(0), count, "Expected 0 rows in the users table after setup")
}

// TestDatabaseMigrations ensures that the migrations run without errors.
func TestDatabaseMigrations(t *testing.T) {
	// Set up the test database with models to test migrations
	SetupGormTestDB(t, &models.User{}, &models.Example{})
	defer TeardownGormTestDB(&models.User{}, &models.Example{})

	// Test that the migrations have been applied successfully
	userTableExists := db.GormDB.Migrator().HasTable(&models.User{})
	assert.True(t, userTableExists, "Expected 'User' table to exist")

	// Optionally, check for other models
	exampleTableExists := db.GormDB.Migrator().HasTable(&models.Example{})
	assert.True(t, exampleTableExists, "Expected 'Example' table to exist")
}

// TestTeardownDatabase ensures that the database tables are dropped after tests.
func TestTeardownDatabase(t *testing.T) {
	// Set up the test database with models to test migrations
	SetupGormTestDB(t, &models.User{}, &models.Example{})
	defer TeardownGormTestDB(&models.User{}, &models.Example{})

	// Drop tables explicitly using Teardown
	TeardownGormTestDB(&models.User{}, &models.Example{})

	// Test that the tables are actually dropped
	userTableExists := db.GormDB.Migrator().HasTable(&models.User{})
	assert.False(t, userTableExists, "Expected 'User' table to be dropped")

	exampleTableExists := db.GormDB.Migrator().HasTable(&models.Example{})
	assert.False(t, exampleTableExists, "Expected 'Example' table to be dropped")
}
