package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xena/helpers"
)

// GettrProfileWebsite returns the website string applied to a Gettr account.
func GettrProfileWebsite(username string) (string, error) {
	website := ""

	request, err := http.NewRequest("GET", "https://api.gettr.com/s/uinf/"+username, nil)
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
	if err != nil {
		return website, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return website, err
	}
	defer response.Body.Close()

	var profile map[string]interface{}

	jsonDecoder := json.NewDecoder(response.Body)
	err = jsonDecoder.Decode(&profile)
	if err != nil {
		return website, err
	}

	data := profile["result"].(map[string]interface{})["data"]
	website = fmt.Sprint(data.(map[string]interface{})["website"])

	return website, nil
}
