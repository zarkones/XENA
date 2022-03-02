package modules

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// GeneratePrivateKey creates a RSA private key.
func GeneratePrivateKey() *rsa.PrivateKey {
	secret, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
	}
	return secret
}

// PublicKeyToPEM converts RSA public key object to PEM encoded string.
func PublicKeyToPEM(publicKey *rsa.PublicKey) string {
	spkiDER, _ := x509.MarshalPKIXPublicKey(publicKey)
	spkiPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: spkiDER,
		},
	)
	return string(spkiPEM)
}

// PrivateKeyToPEM converts RSA private key object to PEM encoded string.
func PrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	spkiDER, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	spkiPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: spkiDER,
		},
	)
	return string(spkiPEM)
}

// ImportPEMPrivateKey converts a PEM encoded private key into the rsa.PublicKey object.
func ImportPEMPrivateKey(spkiPEM string) (*rsa.PrivateKey, error) {
	body, _ := pem.Decode([]byte(spkiPEM))
	privateKey, err := x509.ParsePKCS8PrivateKey(body.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("failed to parse the private key")
	}
	return rsaKey, nil
}

// ImportPEMPublicKey converts a PEM encoded public key into the rsa.PublicKey object.
func ImportPEMPublicKey(spkiPEM string) *rsa.PublicKey {
	body, _ := pem.Decode([]byte(spkiPEM))
	publicKey, err := x509.ParsePKIXPublicKey(body.Bytes)
	if err != nil {
		fmt.Println("Failed to parse public key.")
		return nil
	}
	if publicKey, ok := publicKey.(*rsa.PublicKey); ok {
		return publicKey
	}
	fmt.Println("Failed to parse public key.")
	return nil
}

func EncryptRSAOAEP(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptRSAOAEP(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext)
}

// HashSelf finds the path of the bot executable, reads it and returns MD5 hash of it.
func HashSelf() (string, error) {
	hash := ""

	selfPath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return hash, err
	}

	contentRaw, err := ioutil.ReadFile(selfPath)
	if err != nil {
		fmt.Println(err)
		return hash, err
	}

	md5Hash := md5.Sum(contentRaw)
	hash = hex.EncodeToString(md5Hash[:])

	return hash, nil
}
