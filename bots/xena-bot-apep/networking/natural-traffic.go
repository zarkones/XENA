package networking

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"
	"xena/helpers"
)

type elem struct{ key, val string }

type object []elem

// MarshalJSON converts object type into JSON structure.
func (o object) MarshalJSON() (out []byte, err error) {
	if o == nil {
		return []byte(`null`), nil
	}
	if len(o) == 0 {
		return []byte(`{}`), nil
	}

	out = append(out, '{')
	for _, e := range o {
		key, err := json.Marshal(e.key)
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(e.val)
		if err != nil {
			return nil, err
		}
		out = append(out, key...)
		out = append(out, ':')
		out = append(out, val...)
		out = append(out, ',')
	}
	// replace last ',' with '}'
	out[len(out)-1] = '}'
	return out, nil
}

// SerializedTraffic wraps around naturalTraffic and serializes map to string.
func SerializedTraffic(payload string) (string, error) {
	trafficMap := naturalTraffic(string(payload))
	traffic, err := trafficMap.MarshalJSON()
	return string(traffic), err
}

// naturalTraffic obfuscates 'payload' into JSON-like structure.
func naturalTraffic(payload string) object {
	// Decide on how many keys there will be in the JSON structure.
	indexChar := 0
	maxChars := 126
	minChars := 16

	var jsonObject object

	// Build the JSON structure.
	for indexChar < len(payload) {
		rand.Seed(time.Now().UnixNano())
		chunkSize := rand.Intn(maxChars-minChars) + minChars
		if len(payload) < indexChar+chunkSize {
			chunkSize = len(payload) - indexChar
		}
		key := helpers.RandomPopularWord()
		jsonObject = append(jsonObject, elem{
			key: key,
			val: base64.StdEncoding.EncodeToString([]byte(payload[indexChar : indexChar+chunkSize])),
		})
		indexChar += chunkSize
	}

	return jsonObject
}
