package domains

// Payload bots exchange in pinging each other,
type HandshakeContract struct {
	Id string `json:"id"` // Unique identifier of other bots. (uuid)
}
