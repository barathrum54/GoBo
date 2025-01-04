package main

import (
	"gobo/internal/app"
	"gobo/internal/cache"
	"gobo/internal/db"
	"gobo/internal/logger"
	"gobo/internal/models"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println(".env file loaded successfully.")

	// GORM bağlantısını başlat
	db.ConnectGORM()
	log.Println("Database connection established with GORM.")

	// Migration işlemleri
	models.AutoMigrateExamples(db.GormDB)
	log.Println("Database migrations completed.")

		// Logger yapılandırmasını yükle
	config := logger.DefaultConfig()
	config.Environment = "development" // Geliştirme ortamında çalışıyoruz
	config.OutputPaths = []string{"stdout"} // Sadece terminale yaz

	// Logger'ı başlat
	logger.InitLogger(config)
	// Örnek loglama
	logger.Log.Info("Application started",
		zap.String("service", "gobo"),
		zap.String("status", "running"),
	)
	// Redis bağlantısını başlat
	cache.Connect()

	// Fiber uygulamasını başlat
	application := app.NewApp()

	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
