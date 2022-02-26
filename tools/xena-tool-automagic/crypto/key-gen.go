package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

// KeyPair returns private and public key pair in PEM format.
func KeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "<nil>", "<nil>", err
	}

	return PrivateKeyToPEM(privateKey), PublicKeyToPEM(&privateKey.PublicKey), nil
}
