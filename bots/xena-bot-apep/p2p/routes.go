package p2p

import (
	"net/http"
	"xena/p2p/controllers"
)

type Routing struct {
	peersController controllers.Peers
}

func (routing Routing) Init() {
	http.HandleFunc("/v1/handshake", routing.peersController.Handshake)
}
