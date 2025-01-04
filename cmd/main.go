package main

import (
	"gobo/internal/app"
	"gobo/internal/db"
	"gobo/internal/models"
	"log"

	"github.com/joho/godotenv"
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

	// Fiber uygulamasını başlat
	application := app.NewApp()

	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
