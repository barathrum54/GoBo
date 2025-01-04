package main

import (
	"gobo/internal/app"
	"log"
)

func main() {
	application := app.NewApp()

	// Server'ı başlat
	log.Println("Server is running on http://localhost:3000")
	if err := application.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
