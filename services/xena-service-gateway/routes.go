package main

import "net/http"

func initRoutes() {
	http.HandleFunc("/v1/clients", atilaClientInsert)
	http.HandleFunc("/v1/messages", atilaFetchMessages)
	http.HandleFunc("/v1/messages/ack", atilaMessageAck)
	http.HandleFunc("/v1/services", domenaServiceInsert)
}
