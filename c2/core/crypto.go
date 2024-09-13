package core

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"time"

	cry "github.com/zarkones/xena-crypto"
)

var PrivateKey *rsa.PrivateKey

const PrivKeyPath = "xena.key"

var Certificate *x509.Certificate

const CertPath = "cert.der"

func InitKeys() (err error) {
	// This deferred function would load or create a certificate used for the referse proxy.
	defer func() {
		if err != nil {
			return
		}

		loadExistingCert := func(rawDer *[]byte) {
			privKeyPEM, innerErr := cry.PrivKeyToPEM(PrivateKey)
			if innerErr != nil {
				err = innerErr
				return
			}

			pemBlock := &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: *rawDer,
			}

			pemBytes := pem.EncodeToMemory(pemBlock)

			goproxyCA, innerErr := tls.X509KeyPair(pemBytes, []byte(privKeyPEM))
			if innerErr != nil {
				err = innerErr
				return
			}

			if goproxyCA.Leaf, err = x509.ParseCertificate(goproxyCA.Certificate[0]); err != nil {
				return
			}

			Certificate = goproxyCA.Leaf
		}

		rawDer, innerErr := os.ReadFile(CertPath)
		if innerErr == nil {
			loadExistingCert(&rawDer)
			return
		}

		if !errors.Is(innerErr, os.ErrNotExist) {
			err = innerErr
			return
		}

		Certificate = &x509.Certificate{
			SerialNumber: big.NewInt(2019),
			Subject: pkix.Name{
				Organization: []string{"XENA"},
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(10, 0, 0),
			IsCA:                  true,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
			BasicConstraintsValid: true,
		}

		caBytes, innerErr := x509.CreateCertificate(rand.Reader, Certificate, Certificate, PrivateKey.Public(), PrivateKey)
		if innerErr != nil {
			err = innerErr
			return
		}

		if innerErr := os.WriteFile(CertPath, caBytes, 0777); innerErr != nil {
			err = innerErr
			return
		}
	}()

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
