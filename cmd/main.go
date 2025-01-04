package main

import (
	"gobo/internal/app"
	"gobo/internal/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Veritabanı bağlantısını başlat
	db.Connect()

	// Fiber uygulamasını başlat
	application := app.NewApp()

	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
