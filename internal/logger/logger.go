// Package logger provides utilities for configuring and initializing the application's logger.
// It uses the Uber Zap library for structured and high-performance logging.
package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the global logger instance used throughout the application.
var Log *zap.Logger

// InitLogger initializes the Zap logger with the specified configuration.
// It supports different configurations for "development" and "production" environments.
// If the initialization fails or an invalid environment is provided, the application terminates.
//
// Parameters:
// - config (Config): The logger configuration specifying log level, environment, and output paths.
//
// Behavior:
// - For the "development" environment, a human-readable logging format is used.
// - For the "production" environment, a JSON logging format is used.
// - If the environment is invalid, the function panics.
func InitLogger(config Config) {
	var zapConfig zap.Config

	// Configure the logger based on the specified environment.
	switch config.Environment {
	case "development":
		zapConfig = zap.NewDevelopmentConfig() // Use development-friendly settings.
	case "production":
		zapConfig = zap.NewProductionConfig() // Use production-friendly settings.
	default:
		// Terminate the application if an invalid environment is provided.
		panic("Invalid environment: " + config.Environment)
	}

	// Set the log level (e.g., DEBUG, INFO, ERROR) based on the configuration.
	zapConfig.Level = zap.NewAtomicLevelAt(config.Level)

	// Configure the output paths for the logger (e.g., stdout, files).
	zapConfig.OutputPaths = config.OutputPaths

	// Customize the encoder configuration for log formatting.
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",                 // Key for timestamp in the log.
		LevelKey:       "level",                     // Key for log level.
		NameKey:        "logger",                   // Key for logger name.
		CallerKey:      "caller",                   // Key for caller information.
		MessageKey:     "msg",                      // Key for log message.
		StacktraceKey:  "stacktrace",               // Key for stack traces.
		LineEnding:     zapcore.DefaultLineEnding,  // Line ending character.
		EncodeLevel:    zapcore.CapitalLevelEncoder, // Log level in uppercase (e.g., INFO).
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 time format for timestamp.
		EncodeDuration: zapcore.StringDurationEncoder, // Duration format as string.
		EncodeCaller:   zapcore.ShortCallerEncoder, // Caller information in short format.
	}

	// Build the logger instance using the configured settings.
	var err error
	Log, err = zapConfig.Build()
	if err != nil {
		// Log a fatal error and terminate the application if logger initialization fails.
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Log a message indicating successful logger initialization.
	Log.Info("Logger initialized successfully")
}
