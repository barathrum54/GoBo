// Package logger provides utilities for configuring and initializing the application's logging system.
// It defines the logger configuration structure and default settings.
package logger

import (
	"go.uber.org/zap/zapcore"
)

// Config defines the configuration for the logger.
// It includes options for log level, environment, and output paths.
type Config struct {
	Level       zapcore.Level `json:"level"`       // Log level (e.g., DEBUG, INFO, ERROR)
	Environment string        `json:"environment"` // Logging environment: "development" or "production"
	OutputPaths []string      `json:"outputPaths"` // Output paths for logs (e.g., stdout, file paths)
}

// DefaultConfig returns the default configuration for the logger.
// This default configuration is used if no custom configuration is provided.
// 
// Defaults:
// - Level: zapcore.InfoLevel (logs informational messages and above)
// - Environment: "production" (optimized for production use)
// - OutputPaths: Writes logs to both the terminal (stdout) and a log file (logs/app.log)
//
// Returns:
// - Config: The default logger configuration.
func DefaultConfig() Config {
	return Config{
		Level:       zapcore.InfoLevel,                   // Default to INFO level logging
		Environment: "production",                       // Default to production environment
		OutputPaths: []string{"stdout", "logs/app.log"}, // Log to both stdout and a file
	}
}
