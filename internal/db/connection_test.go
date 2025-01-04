package db_test

import (
	"log"
	"os"
	"testing"

	"gobo/internal/db"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// setupGormTestDB initializes the database connection for testing
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

// teardownGormTestDB cleans up the database after testing
func teardownGormTestDB(t *testing.T) {
	log.Println("[Teardown] Starting test database teardown...")
	sqlDB, err := db.GormDB.DB()
	if err != nil {
		t.Fatalf("[Error] Failed to retrieve SQL DB instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		t.Fatalf("[Error] Failed to close SQL DB connection: %v", err)
	}
	log.Println("[Teardown] Test database teardown completed.")
}

func TestConnectGORM_Success(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB(t)

	// Verify the database connection
	assert.NotNil(t, db.GormDB, "GormDB should not be nil after connection")
	sqlDB, err := db.GormDB.DB()
	assert.NoError(t, err, "Should be able to retrieve SQL DB instance")
	assert.NotNil(t, sqlDB, "SQL DB instance should not be nil")
}

func TestConnectGORM_MissingEnv(t *testing.T) {
	// DATABASE_URL değişkenini temizleyin
	os.Unsetenv("DATABASE_URL")

	// panic'i yakalamak için defer ve recover kullanıyoruz
	defer func() {
		if r := recover(); r != nil {
			t.Logf("[Recovered] Expected panic for missing DATABASE_URL: %v", r)
		} else {
			t.Errorf("Expected a panic for missing DATABASE_URL, but did not get one")
		}
	}()

	// ConnectGORM çağrısı (panic bekleniyor)
	db.ConnectGORM()
}

func TestConnectionPoolSettings(t *testing.T) {
	setupGormTestDB(t)
	defer teardownGormTestDB(t)

	// Retrieve the database connection
	sqlDB, err := db.GormDB.DB()
	assert.NoError(t, err, "Should be able to retrieve SQL DB instance")

	// Check MaxIdleConns
	expectedMaxIdleConns := 10
	actualMaxIdleConns := sqlDB.Stats().Idle // Bu mevcut boşta bağlantı sayısını döndürür.
	assert.LessOrEqual(t, actualMaxIdleConns, expectedMaxIdleConns, "Idle connections should not exceed MaxIdleConns")

	// Check MaxOpenConns
	expectedMaxOpenConns := 100
	// GORM'da aktif bağlantı limitini doğrulama
	assert.LessOrEqual(t, sqlDB.Stats().OpenConnections, expectedMaxOpenConns, "Open connections should not exceed MaxOpenConns")
}

