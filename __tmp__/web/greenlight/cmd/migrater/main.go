package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"greenlight.id/migrations"
)

func main() {
	// Read the database connection string from environment variables
	dbDSN := os.Getenv("GREENLIGHT_DB_DSN")
	if dbDSN == "" {
		log.Fatal("ERROR: GREENLIGHT_DB_DSN environment variable is required")
	}

	// Connect to the database
	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Fatalf("ERROR: could not connect to the database: %v", err)
	}
	defer db.Close()

	// Run the migrations
	if err := runMigrations(db); err != nil {
		log.Fatalf("ERROR: could not run migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully!")
}

func runMigrations(db *sql.DB) error {
	// Assuming you have a function that executes your migrations
	migrationDir := "./migrations"                       // Path to your migration files
	migrator := migrations.NewMigrator(db, migrationDir) // Adjust this to fit your migrations package

	return migrator.Up() // This should be your method to apply the migrations
}
