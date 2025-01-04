package logger_test

import (
	"bytes"
	"os"
	"testing"

	"gobo/internal/logger"

	"go.uber.org/zap/zapcore"
)

func TestInitLogger_DevelopmentConfig(t *testing.T) {
	// Config ayarla
	config := logger.Config{
		Level:       zapcore.DebugLevel,
		Environment: "development",
	}

	// Logger'ın çıktısını yakalamak için bir mock writer oluştur
	tempFile, err := os.CreateTemp("", "testlog")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	config.OutputPaths = []string{tempFile.Name()}

	// Logger'ı başlat
	logger.InitLogger(config)

	// Test bir log yaz
	logger.Log.Debug("Test debug message")
	logger.Log.Info("Test info message")

	// Log dosyasını oku
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	// Log içeriğini kontrol et
	if !bytes.Contains(content, []byte("Test debug message")) {
		t.Errorf("Expected debug message not found in logs")
	}

	if !bytes.Contains(content, []byte("Test info message")) {
		t.Errorf("Expected info message not found in logs")
	}
}

func TestInitLogger_ProductionConfig(t *testing.T) {
	// Config ayarla
	config := logger.Config{
		Level:       zapcore.InfoLevel,
		Environment: "production",
		OutputPaths: []string{"stdout"},
	}

	// Logger'ı başlat
	logger.InitLogger(config)

	// Logger'ın doğru şekilde başlatıldığını kontrol et
	if logger.Log == nil {
		t.Fatalf("Logger was not initialized")
	}

	logger.Log.Info("Test info message")
}

func TestInitLogger_InvalidConfig(t *testing.T) {
    config := logger.Config{
        Level:       zapcore.InfoLevel,
        Environment: "invalid", // Geçersiz ortam
        OutputPaths: []string{"stdout"},
    }

    defer func() {
        if r := recover(); r != nil {
            t.Log("Panic caught as expected")
        } else {
            t.Errorf("Expected panic for invalid environment, but did not get one")
        }
    }()

    logger.InitLogger(config)
}


func TestDefaultConfig(t *testing.T) {
	// Default config'i al
	defaultConfig := logger.DefaultConfig()

	// Default config değerlerini kontrol et
	if defaultConfig.Level != zapcore.InfoLevel {
		t.Errorf("Expected default level to be InfoLevel, got %v", defaultConfig.Level)
	}

	if defaultConfig.Environment != "production" {
		t.Errorf("Expected default environment to be 'production', got %v", defaultConfig.Environment)
	}

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
