// Package models contains the application's database models and related functionality.
// This file defines the User model and its associated migration logic.
package models

import "gorm.io/gorm"

// User represents the "users" table in the database.
// Fields:
// - ID: The primary key of the record.
// - Username: A required string field with a maximum length of 100 characters, must be unique.
// - Password: A required string field for storing the user's password.
// - Email: A required string field with a maximum length of 100 characters, must be unique.
type User struct {
	ID       uint   `gorm:"primaryKey"`                    // Primary key for the record.
	Username string `gorm:"type:varchar(100);unique;not null"` // Username field, unique and required with max length of 100 characters.
	Password string `gorm:"type:varchar(100);not null"`      // Password field, required with max length of 100 characters.
	Email    string `gorm:"type:varchar(100);unique;not null"` // Email field, unique and required with max length of 100 characters.
}

// AutoMigrateUsers ensures the "users" table schema is up to date.
// It uses GORM's AutoMigrate feature to create or update the table as needed.
//
// Parameters:
// - db (*gorm.DB): The GORM database connection instance.
//
// Returns:
// - error: Returns an error if the migration fails.
//
// Behavior:
// - If the migration fails, the function panics with a detailed error message.
// - On success, the function completes without error.
func AutoMigrateUsers(db *gorm.DB) error {
	// Perform the migration for the User model.
	err := db.AutoMigrate(&User{})
	if err != nil {
		// Panic with a detailed error message if the migration fails.
		panic("Failed to migrate user table: " + err.Error())
	}

	// Return nil if the migration is successful.
	return nil
}
