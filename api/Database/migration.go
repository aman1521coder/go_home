package database

import (
	"database/sql"
	"fmt"
	"log"
)

func RunMigration(db *sql.DB) {
	migrations := []string{
		createUsersTable,
		createItemsTable,
		createItemsUserIndex,
	}
	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			log.Fatalf("Error running migration: %v", err)
		}
	}
	fmt.Println("Migrations completed successfully")
}

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	username VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	is_admin BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

const createItemsTable = `
CREATE TABLE IF NOT EXISTS items (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	price DECIMAL(10, 2) NOT NULL,
	selling_price DECIMAL(10, 2) NOT NULL,
	image VARCHAR(500),
	quantity INTEGER DEFAULT 0,
	is_sold BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`
const createItemsUserIndex = `
CREATE INDEX IF NOT EXISTS idx_items_user_id ON items(user_id);`
