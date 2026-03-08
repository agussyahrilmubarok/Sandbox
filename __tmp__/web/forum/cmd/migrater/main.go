package main

import (
	"fmt"
	"log"
	"os"

	"forum.id/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.LoadViper()

	host := cfg.GetString("POSTGRES.HOST")
	port := cfg.GetInt("POSTGRES.PORT")
	user := cfg.GetString("POSTGRES.USER")
	password := cfg.GetString("POSTGRES.PASSWORD")
	name := cfg.GetString("POSTGRES.NAME")
	sslMode := cfg.GetString("POSTGRES.SSL_MODE")
	maxOpenConns := cfg.GetInt("POSTGRES.MAX_OPEN_CONNS")
	maxIdleConns := cfg.GetInt("POSTGRES.MAX_IDLE_CONNS")
	maxIdleTime := cfg.GetDuration("POSTGRES.MAX_IDLE_TIME")

	fmt.Printf("Parsed config: host=%s, port=%d, user=%s, password=%s, name=%s, sslMode=%s, maxOpenConns=%d, maxIdleConns=%d, maxIdleTime=%v\n",
		host, port, user, password, name, sslMode, maxOpenConns, maxIdleConns, maxIdleTime)

	m, err := postgresMigrate(user, password, host, port, name, sslMode)
	if err != nil {
		log.Fatal("failed to create migration instance:", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("failed to rollback migration:", err)
		}
		log.Println("migration rolled back successfully!")
	} else {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("failed to apply migration:", err)
		}
		log.Println("migration applied successfully!")
	}
}

func postgresMigrate(user, password, host string, port int, name, sslMode string) (*migrate.Migrate, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user, password, host, port, name, sslMode)

	fmt.Printf("DSN: %s\n", dsn)

	m, err := migrate.New(
		"file://db/migrations",
		dsn,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance: %w", err)
	}

	return m, nil
}
