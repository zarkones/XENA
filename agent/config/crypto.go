package config

import (
	"crypto/rsa"
	"os"

	cry "github.com/zarkones/xena-crypto"
)

var PrivateKey *rsa.PrivateKey

func InitKeys(privKeyPath string) (err error) {
	rawPrivKey, err := os.ReadFile(privKeyPath)
	if err != nil {
		return firstTimeInit(privKeyPath)
	}

	PrivateKey, err = cry.ImportPrivKeyPEM(rawPrivKey)
	if err != nil {
		return err
	}

	return nil
}

func firstTimeInit(privKeyPath string) (err error) {
	PrivateKey, err = cry.GenPrivKey()
	if err != nil {
		return err
	}
	privKeyPEM, err := cry.PrivKeyToPEM(PrivateKey)
	if err != nil {
		return err
	}
	if err := os.WriteFile(privKeyPath, []byte(privKeyPEM), 0777); err != nil {
		return err
	}
	return nil
}
