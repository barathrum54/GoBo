// Package db provides utilities for database connection and configuration using GORM.
// It supports connecting to a PostgreSQL database and setting connection pool parameters.
package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormDB is the global GORM database instance used for interacting with the PostgreSQL database.
var GormDB *gorm.DB

// ConnectGORM initializes a GORM connection to the PostgreSQL database.
// It retrieves the database connection string (DSN) from the DATABASE_URL environment variable,
// establishes the connection, and configures the connection pool.
//
// If the DATABASE_URL environment variable is not set or the connection fails, the application will terminate.
func ConnectGORM() {
	// Retrieve the database connection string from the environment variable.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Terminate the application if the DATABASE_URL is not set.
		panic("DATABASE_URL environment variable is not set")
	}

	// Open a GORM connection using the PostgreSQL driver and default logger.
	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Use the default logger in Info mode
	})
	if err != nil {
		// Log a fatal error and terminate if the connection fails.
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Retrieve the underlying SQL database instance to configure connection pool settings.
	sqlDB, err := GormDB.DB()
	if err != nil {
		// Log a fatal error and terminate if the SQL database instance cannot be retrieved.
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Configure the connection pool:
	// - SetMaxIdleConns: Maximum number of idle connections in the pool.
	// - SetMaxOpenConns: Maximum number of open connections to the database.
	// - SetConnMaxLifetime: Maximum amount of time a connection may be reused.
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Log a message indicating a successful connection.
	log.Println("Connected to the database using GORM!")
}
