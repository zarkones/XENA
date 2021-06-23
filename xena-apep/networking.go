package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/* Returns a new port which should be used for binding onto the network. Day of the month is set as the seed. */
func currentActivePort() int {
	rand.Seed(int64(time.Now().Day()))
	return rand.Intn(60666-3003) + 3003
}

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
