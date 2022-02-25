package domains

import (
	"encoding/json"
)

// Payload for endpoint of Atila for message's ack.
type MessageAck struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type Message struct {
	Id      string `json:"id"`      // Unique identifier.
	From    string `json:"from"`    // Node which originally issued the message.
	To      string `json:"to"`      // Which node should receive message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	Status  string `json:"status"`  // Message's state.
	ReplyTo string `json:"replyTo"` // Original message ID.
}

// String serializes the message domain into a JSON string.
func (message *Message) String() (string, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return "<nil>", err
	}
	return string(jsonMessage), nil
}
