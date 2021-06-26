package xena

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

/* Entity storage. */
var entities = make(map[string]Entity)
