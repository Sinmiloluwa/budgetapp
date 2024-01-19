package main

import "database/sql"

func createUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL`
}