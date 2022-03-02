package models

// Model of 'details' table.
type Details struct {
	Id         string `json:"id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
