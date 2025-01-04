package logger

import (
	"go.uber.org/zap/zapcore"
)

// Config defines the configuration for the logger
type Config struct {
	Level       zapcore.Level `json:"level"`       // Log seviyesi (DEBUG, INFO, ERROR)
	Environment string        `json:"environment"` // Ortam: "development" veya "production"
	OutputPaths []string      `json:"outputPaths"` // Logların yazılacağı dosyalar veya stdout
}

// DefaultConfig returns the default configuration for the logger
func DefaultConfig() Config {
	return Config{
		Level:       zapcore.InfoLevel,
		Environment: "production",
		OutputPaths: []string{"stdout", "logs/app.log"}, // Terminale ve dosyaya yaz
	}
}
