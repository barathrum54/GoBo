// Package logger provides utilities for configuring and initializing the application's logger.
// It uses the Uber Zap library for structured and high-performance logging and integrates with Sentry for error tracking.
package logger

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the global logger instance used throughout the application.
var Log *zap.Logger

// InitLogger initializes the Zap logger with the specified configuration and Sentry integration.
// It supports different configurations for "development" and "production" environments.
// If the initialization fails or an invalid environment is provided, the application terminates.
//
// Parameters:
// - config (Config): The logger configuration specifying log level, environment, output paths, and Sentry DSN.
//
// Behavior:
// - For the "development" environment, a human-readable logging format is used.
// - For the "production" environment, a JSON logging format is used.
// - If a Sentry DSN is provided, errors and higher-severity logs are sent to Sentry.
func InitLogger(config Config) {
	// Initialize Sentry if a DSN is provided.
	if config.SentryDSN != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         config.SentryDSN,
			Environment: config.Environment,
		})
		if err != nil {
			log.Fatalf("Failed to initialize Sentry: %v", err)
		}
		// Ensure Sentry flushes events before exiting the application.
		defer sentry.Flush(2 * time.Second)
	}

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

	// Build the logger instance using the configured settings.
	var err error
	Log, err = zapConfig.Build(zap.Hooks(sentryHook))
	if err != nil {
		// Log a fatal error and terminate the application if logger initialization fails.
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Log a message indicating successful logger initialization.
	Log.Info("Logger initialized successfully")
}

// sentryHook is a Zap hook that sends error-level and above logs to Sentry.
func sentryHook(entry zapcore.Entry) error {
	if entry.Level >= zapcore.ErrorLevel {
		// Capture the error-level log in Sentry.
		sentry.CaptureMessage(entry.Message)
	}
	return nil
}
