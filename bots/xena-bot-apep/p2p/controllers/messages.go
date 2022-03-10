package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xena/config"
	"xena/domains"
	"xena/gateway"
)

type Messages struct{}

func (messagesController *Messages) Execute(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "*")
	response.Header().Set("Access-Control-Max-Age", "86400")

	fmt.Println(request.Method == http.MethodOptions)
	fmt.Println(request.Method + http.MethodOptions)

	if request.Method == http.MethodOptions {
		return
	}

	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	var message domains.Message

	err := jsonDecoder.Decode(&message)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, err := gateway.InterpretMessage(message)
	reply.To = config.ID
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	serializedMessage, err := reply.String()
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(serializedMessage))
}
