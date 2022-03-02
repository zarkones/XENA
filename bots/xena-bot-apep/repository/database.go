package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Database connection.
type Database struct {
	name  string
	Query *sql.DB
}

var DB Database = Database{}

// Init creates the database file and runs the migration.
func (db *Database) Init(name string) error {
	db.name = name

	query, err := sql.Open("sqlite3", "./"+db.name+".db")
	if err != nil {
		return err
	}
	db.Query = query

	return nil
}

// RunMigrations creates the database's tables, if they don't exist already.
func (db *Database) RunMigrations() error {
	statement, err := db.Query.Prepare(`
		CREATE TABLE
		IF NOT EXISTS
		details (
			id TEXT,
			private_key TEXT,
			public_key TEXT
		)
	`)
	if err != nil {
		return err
	}

	statement.Exec()

	return nil
}
