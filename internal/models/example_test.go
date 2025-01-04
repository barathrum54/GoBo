package models

import (
	"gobo/internal/db"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the test database using GORM
func setupGormTestDB() {
	log.Println("[Setup] Starting GORM test database setup...")

	db.ConnectGORM()

	// Tabloları oluştur
	AutoMigrateExamples(db.GormDB)

	log.Println("[Setup] Test database setup completed.")
}

// teardownGormTestDB drops all test tables
func teardownGormTestDB() {
	log.Println("[Teardown] Starting GORM test database teardown...")

	sqlDB, _ := db.GormDB.DB()
	sqlDB.Exec("DROP TABLE IF EXISTS examples")

	log.Println("[Teardown] Test database teardown completed.")
}

func TestCreateExampleGorm(t *testing.T) {
	setupGormTestDB()
	defer teardownGormTestDB()

	// Yeni bir kayıt ekle
	example := Example{Name: "Test Name"}
	result := db.GormDB.Create(&example)
	assert.NoError(t, result.Error, "Error occurred while creating an example")

	// Veritabanını kontrol et
	var count int64
	db.GormDB.Model(&Example{}).Count(&count)
	assert.Equal(t, int64(1), count, "Expected 1 row in examples table")
}

func TestGetExamplesGorm(t *testing.T) {
	setupGormTestDB()
	defer teardownGormTestDB()

	// Test verileri ekle
	db.GormDB.Create(&Example{Name: "Example 1"})
	db.GormDB.Create(&Example{Name: "Example 2"})

	// Tablodan kayıtları getir
	var examples []Example
	result := db.GormDB.Find(&examples)
	assert.NoError(t, result.Error, "Error occurred while retrieving examples")

	// Kayıt sayısını doğrula
	assert.Equal(t, 2, len(examples), "Expected 2 examples")
	assert.Equal(t, "Example 1", examples[0].Name)
	assert.Equal(t, "Example 2", examples[1].Name)
}
