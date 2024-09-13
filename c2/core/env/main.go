package env

import (
	"errors"
	"os"
)

var (
	AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
	PROXY_HOST = os.Getenv("PROXY_HOST")
	PROXY_PORT = os.Getenv("PROXY_PORT")
)

func Validate() (err error) {
	if len(AUTH_TOKEN) == 0 {
		return errors.New("AUTH_TOKEN not supplied")
	}
	if len(PROXY_HOST) == 0 {
		PROXY_HOST = "127.0.0.1"
	}
	if len(PROXY_PORT) == 0 {
		PROXY_PORT = "8989"
	}
	return nil
}
