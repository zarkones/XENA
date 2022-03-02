package repository

import (
	"fmt"
	"xena/repository/models"
)

type detailsRepo struct {
}

var DetailsRepo detailsRepo = detailsRepo{}

// Insert saves basic information about the bot so that it can maintain session with other services and peers.
func (details detailsRepo) Insert(id, privateKey, publicKey string) error {
	statement, err := DB.Query.Prepare(`
		INSERT INTO
		details (
			id,
			private_key,
			public_key
		) VALUES (?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
		return err
	}

	statement.Exec(id, privateKey, publicKey)

	return nil
}

// Get returns details about the bot, such are: id, private key and public key.
func (details detailsRepo) Get() (models.Details, error) {
	maybeDetails := models.Details{}

	rows, err := DB.Query.Query(`
		SELECT
			id,
			private_key,
			public_key
		FROM details
		LIMIT 1
	`)
	if err != nil {
		return maybeDetails, nil
	}

	for rows.Next() {
		rows.Scan(&maybeDetails.Id, &maybeDetails.PrivateKey, &maybeDetails.PublicKey)
	}

	return maybeDetails, nil
}
