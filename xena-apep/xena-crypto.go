package xena

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

/* Generates a private key. */
func generatePrivateKey() *rsa.PrivateKey {
	secret, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
	}
	return secret
}

func encryptRSAOAEP(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func decryptRSAOAEP(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext)
}
