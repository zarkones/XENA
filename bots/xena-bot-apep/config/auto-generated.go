package config

import "crypto/rsa"

// Generate the unique bot identifier.
var ID string

// Key-pair used for signing and verifying messages.
var PrivateIdentificationKey *rsa.PrivateKey
var PublicIdentificationKey *rsa.PublicKey
