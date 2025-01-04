package models

import "gorm.io/gorm"

// Example represents the example table
type Example struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}

// AutoMigrateExamples migrates the example table
func AutoMigrateExamples(db *gorm.DB) error {
	err := db.AutoMigrate(&Example{})
	if err != nil {
		panic("Failed to migrate example table: " + err.Error())
	}
	return nil


}
