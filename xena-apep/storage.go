package main

/* Internet Address. x, y, z, w; make a complate internet addres. Examples: 127.0.0.1, 168.168.1.1, 255.255.255.255, etc... */
type InternetAddress struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
	Z uint8 `json:"z"`
	W uint8 `json:"w"`
}

/* Internet Port */
type InternetPort uint16

/* Entity. */
type Entity struct {
	Id        string          `json:"id"`              // Unique identifier.
	Address   InternetAddress `json:"internetAddress"` // Internet address of node.
	CreatedAt int64           `json:"createdAt"`       // Time object of when this entity was created.
}

/* Message. */
type Message struct {
	Id      string `json:"id"`      // Unique identifier.
	From    string `json:"from"`    // Node which originally issued the message.
	To      string `json:"to"`      // Which node should receive message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	Status  string `json:"status"`  // Message's state.
	ReplyTo string `json:"replyTo"` // Original message ID.
}

/* Entity storage. */
var entities = make(map[string]Entity)
var messages = make(map[string]Message)
