package p2p

import (
	"net/http"
)

type P2P struct {
	router Routing
}

// BootServer ignites a HTTP server used for Peer 2 Peer communication.
func (p2p *P2P) BootServer(port string) error {
	p2p.router.Init()

	// Listen on the port.
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}

	return nil
}
