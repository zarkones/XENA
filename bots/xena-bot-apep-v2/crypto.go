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

// importPEMPublicKey converts a PEM encoded public key into the rsa.PublicKey object.
func importPEMPublicKey(spkiPEM string) *rsa.PublicKey {
	body, _ := pem.Decode([]byte(spkiPEM))
	publicKey, _ := x509.ParsePKIXPublicKey(body.Bytes)
	if publicKey, ok := publicKey.(*rsa.PublicKey); ok {
		return publicKey
	}
	return nil
}
