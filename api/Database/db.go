package database

import (
	"database/sql"
	"fmt"
	"log"

	"primeauction/api/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	cfg, err := config.Loadconfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	connStr := cfg.GetDBConnectionString()

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	RunMigration(DB)

	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
