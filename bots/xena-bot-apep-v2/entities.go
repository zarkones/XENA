package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload bots exchange in pinging each other,
type PingContract struct {
	Id string `json:"id"` // Unique identifier of other bots. (uuid)
}

// Port property of the IP protocol.
type InternetPort uint16

// Entity. (a bot peer)
type Entity struct {
	Id      string `json:"id"`      // Unique identifier of other bots. (uuid)
	Address string `json:"address"` // Internet protocol address.
}

// In memory storage of entities. (bot peers)
var entitiesPool []Entity

// In memory storage of messages which are meant for different receivers.
var messagePoolOfOthers []Message

// messageHandler deals with incoming HTTP requests.
func messageHandler(response http.ResponseWriter, request *http.Request) {
	// Read and prepare the body.
	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	var message Message

	err := jsonDecoder.Decode(&message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// If message is meant for this bot, interpret and issue a reply,
	// if not then store it in the pool of messages meat for different bots.
	if message.To == id.String() {
		reply, err := interpretMessage(message)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = sendMessage(reply)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = messageAck(reply.ReplyTo)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		messagePoolOfOthers = append(messagePoolOfOthers, message)
	}
}

// pingHandler deals with pings from other bot peers.
func pingHandler(response http.ResponseWriter, request *http.Request) {
	// Read and prepare the body.
	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	var ping PingContract

	err := jsonDecoder.Decode(&ping)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

// bootServer ignites a HTTP server used for Peer 2 Peer communication.
func bootServer() {
	// Routes.
	http.HandleFunc("/v1/messages", messageHandler)
	http.HandleFunc("/v1/ping", pingHandler)

	// Listen on the port.
	err := http.ListenAndServe(":"+peerPort, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
