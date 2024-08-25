package config

import "crypto/rsa"

// Main loop configuration.
var MaxLoopWait int = 6
var MinLoopWait int = 3

// Command & control server.
var TrustedPubKey *rsa.PublicKey
var GatewayHost string = "http://127.0.0.1:8080"
var PubSubEnabled bool = false

// Persistence
var PersistAtPath string
var PersistAtRegKey string

var AgentID = ""
var PathToAgentID string
var PathToPrivKey string
