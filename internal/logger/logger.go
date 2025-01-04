package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger initializes the Zap logger with the given configuration
func InitLogger(config Config) {
	var zapConfig zap.Config

	if config.Environment == "development" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	// Log seviyesini ayarla
	zapConfig.Level = zap.NewAtomicLevelAt(config.Level)

	// Log çıktı yollarını ayarla
	zapConfig.OutputPaths = config.OutputPaths
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Logger'ı başlat
	var err error
	Log, err = zapConfig.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer Log.Sync() // Buffer temizliği

	Log.Info("Logger initialized successfully")
}
