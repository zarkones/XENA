package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xena/p2p/domains"
)

type Peers struct{}

func (controller *Peers) Handshake(response http.ResponseWriter, request *http.Request) {
	// Read and prepare the body.
	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	var handshake domains.HandshakeContract

	err := jsonDecoder.Decode(&handshake)
	if err != nil {
		fmt.Println(err)
		return
	}
}
