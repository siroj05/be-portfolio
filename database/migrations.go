package database

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// RunMigrations runs goose migrations from the embedded SQL files
func RunMigrations(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("goose run migrations: failed to set dialect: %v", err)
	}

	log.Println("goose run migrations: checking and running database migrations...")
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("goose run migrations: failed to run migrations: %v", err)
	}
	log.Println("goose run migrations: database migrations completed successfully!")
}
