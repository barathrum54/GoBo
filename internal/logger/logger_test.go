// Package logger_test contains tests for the logger package.
// These tests validate logger initialization, configuration, and default behavior.
package logger_test

import (
	"bytes"
	"os"
	"testing"

	"gobo/internal/logger"

	"go.uber.org/zap/zapcore"
)

// TestInitLogger_DevelopmentConfig verifies that the logger initializes correctly
// in the "development" environment with debug-level logging.
func TestInitLogger_DevelopmentConfig(t *testing.T) {
	// Create a logger configuration for the development environment.
	config := logger.Config{
		Level:       zapcore.DebugLevel,  // Set log level to DEBUG
		Environment: "development",      // Set environment to development
	}

	// Create a temporary file to capture logger output for testing.
	tempFile, err := os.CreateTemp("", "testlog")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file after the test.

	config.OutputPaths = []string{tempFile.Name()} // Direct logger output to the temporary file.

	// Initialize the logger with the specified configuration.
	logger.InitLogger(config)

	// Log test messages to verify output.
	logger.Log.Debug("Test debug message") // DEBUG level message
	logger.Log.Info("Test info message")   // INFO level message

	// Read the contents of the temporary log file.
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	// Verify that the log contains the debug message.
	if !bytes.Contains(content, []byte("Test debug message")) {
		t.Errorf("Expected debug message not found in logs")
	}

	// Verify that the log contains the info message.
	if !bytes.Contains(content, []byte("Test info message")) {
		t.Errorf("Expected info message not found in logs")
	}
}

// TestInitLogger_ProductionConfig verifies that the logger initializes correctly
// in the "production" environment with info-level logging.
func TestInitLogger_ProductionConfig(t *testing.T) {
	// Create a logger configuration for the production environment.
	config := logger.Config{
		Level:       zapcore.InfoLevel,   // Set log level to INFO
		Environment: "production",       // Set environment to production
		OutputPaths: []string{"stdout"}, // Log output to standard output
	}

	// Initialize the logger with the specified configuration.
	logger.InitLogger(config)

	// Verify that the logger instance is initialized.
	if logger.Log == nil {
		t.Fatalf("Logger was not initialized")
	}

	// Log a test message to verify output.
	logger.Log.Info("Test info message") // INFO level message
}

// TestInitLogger_InvalidConfig tests the behavior of the logger when an invalid environment is provided.
// It expects the logger initialization to panic.
func TestInitLogger_InvalidConfig(t *testing.T) {
	config := logger.Config{
		Level:       zapcore.InfoLevel,   // Set log level to INFO
		Environment: "invalid",          // Invalid environment
		OutputPaths: []string{"stdout"}, // Log output to standard output
	}

	// Use defer to recover from the expected panic.
	defer func() {
		if r := recover(); r != nil {
			t.Log("Panic caught as expected") // Log the recovered panic for debugging.
		} else {
			t.Errorf("Expected panic for invalid environment, but did not get one")
		}
	}()

	// Attempt to initialize the logger with an invalid configuration.
	logger.InitLogger(config)
}

// TestDefaultConfig verifies the default logger configuration values.
func TestDefaultConfig(t *testing.T) {
	// Retrieve the default logger configuration.
	defaultConfig := logger.DefaultConfig()

	// Validate the default log level.
	if defaultConfig.Level != zapcore.InfoLevel {
		t.Errorf("Expected default level to be InfoLevel, got %v", defaultConfig.Level)
	}

	// Validate the default environment setting.
	if defaultConfig.Environment != "production" {
		t.Errorf("Expected default environment to be 'production', got %v", defaultConfig.Environment)
	}

	// Validate the default output paths.
	expectedOutputs := []string{"stdout", "logs/app.log"}
	if len(defaultConfig.OutputPaths) != len(expectedOutputs) {
		t.Errorf("Expected default output paths to have length %d, got %d", len(expectedOutputs), len(defaultConfig.OutputPaths))
	}
	for i, path := range expectedOutputs {
		if defaultConfig.OutputPaths[i] != path {
			t.Errorf("Expected output path %s, got %s", path, defaultConfig.OutputPaths[i])
		}
	}
}
