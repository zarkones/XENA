package core

import (
	"crypto/rsa"
	"os"

	cry "github.com/zarkones/xena-crypto"
)

var PrivateKey *rsa.PrivateKey

var PrivKeyPath = "xena.key"

func InitKeys() (err error) {
	rawPrivKey, err := os.ReadFile(PrivKeyPath)
	if err != nil {
		return firstTimeInit()
	}

	PrivateKey, err = cry.ImportPrivKeyPEM(rawPrivKey)
	if err != nil {
		return err
	}

	return nil
}

func firstTimeInit() (err error) {
	PrivateKey, err = cry.GenPrivKey()
	if err != nil {
		return err
	}
	privKeyPEM, err := cry.PrivKeyToPEM(PrivateKey)
	if err != nil {
		return err
	}
	if err := os.WriteFile(PrivKeyPath, []byte(privKeyPEM), 0777); err != nil {
		return err
	}
	return nil
}
