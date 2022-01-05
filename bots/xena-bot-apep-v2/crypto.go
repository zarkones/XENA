package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// generatePrivateKey creates a RSA private key.
func generatePrivateKey() *rsa.PrivateKey {
	secret, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
	}
	return secret
}

// publicKeyToPEM converts RSA public key object to PEM encoded string.
func publicKeyToPEM(publicKey *rsa.PublicKey) string {
	spkiDER, _ := x509.MarshalPKIXPublicKey(publicKey)
	spkiPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PULIC KEY",
			Bytes: spkiDER,
		},
	)
	return string(spkiPEM)
}
