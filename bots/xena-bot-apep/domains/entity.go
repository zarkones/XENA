package domains

// Entity. (a bot peer)
type Entity struct {
	Id      string `json:"id"`      // Unique identifier of other bots. (uuid)
	Address string `json:"address"` // Internet protocol address.
}
