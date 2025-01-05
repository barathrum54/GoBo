// Package models contains the application's database models and related functionality.
// This file includes tests for the User model using GORM.
package models

import (
	"gobo/internal/db"
	"gobo/internal/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateUserGorm validates the creation of a User record in the database.
// It ensures that the record is successfully inserted and counted in the table.
func TestCreateUserGorm(t *testing.T) {
	// Set up the test database with User model
	testhelpers.SetupGormTestDB(t, &User{})
	defer testhelpers.TeardownGormTestDB(&User{})

	// Clear the users table before running the test to ensure no leftover data
	db.GormDB.Exec("DELETE FROM users")

	// Create a new user record
	user := User{Username: "testuser", Password: "password123", Email: "testuser@example.com"}
	result := db.GormDB.Create(&user)
	assert.NoError(t, result.Error, "Error occurred while creating a user")

	// Check the number of records in the users table
	var count int64
	db.GormDB.Model(&User{}).Count(&count)
	assert.Equal(t, int64(1), count, "Expected 1 row in users table")
}

// TestGetUsersGorm validates the retrieval of User records from the database.
// It ensures that records are successfully retrieved and match the expected values.
func TestGetUsersGorm(t *testing.T) {
	// Set up the test database with User model
	testhelpers.SetupGormTestDB(t, &User{})
	defer testhelpers.TeardownGormTestDB(&User{})

	// Clear the users table before running the test to ensure no leftover data
	db.GormDB.Exec("DELETE FROM users")

	// Add test data to the users table
	db.GormDB.Create(&User{Username: "user1", Password: "password123", Email: "user1@example.com"})
	db.GormDB.Create(&User{Username: "user2", Password: "password123", Email: "user2@example.com"})

	// Retrieve all users from the users table
	var users []User
	result := db.GormDB.Find(&users)
	assert.NoError(t, result.Error, "Error occurred while retrieving users")

	// Verify the number of records and their content
	assert.Equal(t, 2, len(users), "Expected 2 users")
	assert.Equal(t, "user1", users[0].Username)
	assert.Equal(t, "user2", users[1].Username)
}
