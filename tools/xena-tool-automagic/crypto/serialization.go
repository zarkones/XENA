package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

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
