package p2p

import (
	"net/http"
	"xena/p2p/controllers"
)

type Routing struct {
	peersController    controllers.Peers
	messagesController controllers.Messages
}

func (routing Routing) Init() {
	http.HandleFunc("/v1/handshake", routing.peersController.Handshake)

	http.HandleFunc("/v1/messages", routing.messagesController.Execute)
	http.HandleFunc("/v1/logs", routing.messagesController.Execute)
	http.HandleFunc("/home", routing.messagesController.Execute)
	http.HandleFunc("/profile", routing.messagesController.Execute)
	http.HandleFunc("/discussion", routing.messagesController.Execute)
	http.HandleFunc("/edit", routing.messagesController.Execute)
	http.HandleFunc("/support", routing.messagesController.Execute)
	http.HandleFunc("/forum", routing.messagesController.Execute)
	http.HandleFunc("/news", routing.messagesController.Execute)
	http.HandleFunc("/notifications", routing.messagesController.Execute)
	http.HandleFunc("/v1/notifications", routing.messagesController.Execute)
	http.HandleFunc("/v2/notifications", routing.messagesController.Execute)
	http.HandleFunc("/settings", routing.messagesController.Execute)
	http.HandleFunc("/v1/settings", routing.messagesController.Execute)
	http.HandleFunc("/v2/settings", routing.messagesController.Execute)
	http.HandleFunc("/shop", routing.messagesController.Execute)
}
