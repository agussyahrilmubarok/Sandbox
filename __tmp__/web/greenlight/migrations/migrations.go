package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Migrator struct {
	db           *sql.DB
	migrationDir string
}

func NewMigrator(db *sql.DB, migrationDir string) *Migrator {
	return &Migrator{db: db, migrationDir: migrationDir}
}

// Up applies all migrations in the migration directory
func (m *Migrator) Up() error {
	files, err := ioutil.ReadDir(m.migrationDir)
	if err != nil {
		return fmt.Errorf("could not read migration directory: %w", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			if err := m.applyMigration(file.Name()); err != nil {
				return err
			}
		}
	}
	return nil
}

// applyMigration executes a single migration
func (m *Migrator) applyMigration(filename string) error {
	filePath := filepath.Join(m.migrationDir, filename)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read migration file %s: %w", filename, err)
	}

	_, err = m.db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("could not apply migration %s: %w", filename, err)
	}

	fmt.Printf("Applied migration: %s\n", filename)
	return nil
}
