package main

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"
)

// serializedTraffic wraps around naturalTraffic and serializes map to string.
func serializedTraffic(payload string) (string, error) {
	trafficMap := naturalTraffic(string(payload))
	traffic, err := json.Marshal(trafficMap)
	return string(traffic), err
}

// naturalTraffic obfuscates 'payload' into JSON-like structure.
func naturalTraffic(payload string) map[string]string {
	// Decide on how many keys there will be in the JSON structure.
	indexChar := 0
	maxChars := 126
	minChars := 16

	var jsonObject = make(map[string]string)

	// Build the JSON structure.
	for indexChar < len(payload) {
		rand.Seed(time.Now().UnixNano())
		chunkSize := rand.Intn(maxChars-minChars) + minChars
		if len(payload) < indexChar+chunkSize {
			chunkSize = len(payload) - indexChar
		}
		key := randomPopularWord()
		jsonObject[key] = base64.StdEncoding.EncodeToString([]byte(payload[indexChar : indexChar+chunkSize]))
		indexChar += chunkSize
	}

	return jsonObject
}
