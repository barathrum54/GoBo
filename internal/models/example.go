// Package models contains the application's database models and related functionality.
// This file defines the Example model and its associated migration logic.
package models

import "gorm.io/gorm"

// Example represents the "examples" table in the database.
// Fields:
// - ID: The primary key of the record.
// - Name: A required string field with a maximum length of 100 characters.
type Example struct {
	ID   uint   `gorm:"primaryKey"`                    // Primary key for the record.
	Name string `gorm:"type:varchar(100);not null"`    // Name field, required with a max length of 100 characters.
}

// AutoMigrateExamples ensures the "examples" table schema is up to date.
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
func AutoMigrateExamples(db *gorm.DB) error {
	// Perform the migration for the Example model.
	err := db.AutoMigrate(&Example{})
	if err != nil {
		// Panic with a detailed error message if the migration fails.
		panic("Failed to migrate example table: " + err.Error())
	}

	// Return nil if the migration is successful.
	return nil
}
