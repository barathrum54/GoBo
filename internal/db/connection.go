package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

// Connect initializes a connection to the PostgreSQL database
func Connect(optionalDSN ...string) {
	dsn := os.Getenv("DATABASE_URL")
	if len(optionalDSN) > 0 {
		dsn = optionalDSN[0]
	}
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Conn, err = pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL!")
}

