package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDb() (*sql.DB, error) {
	// it creates a sql.DB object that acts as a connection pool.
	// The database connections are automatically managed behind the scenes.
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// check if the connection establish
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// read schema design database from .sql file
	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("error reading schema.sql: %v", err)
	}

	// creating tables
	if _, err := db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("error executing schema.sql: %v", err)
	}

	return db, nil
}
