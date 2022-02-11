package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func atilaClientInsert(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	nextReq, err := http.NewRequest("POST", atilaHost+"/v1/clients", bytes.NewBuffer(body))
	nextReq.Header.Set("Content-Type", "application/json")
	nextReq.Header.Set("User-Agent", request.UserAgent())
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	client := &http.Client{}

	nextResp, err := client.Do(nextReq)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	defer nextResp.Body.Close()

	respBody, err := ioutil.ReadAll(nextResp.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for the correct status code.
	if nextResp.StatusCode != http.StatusCreated && nextResp.StatusCode != http.StatusConflict {
		fmt.Println("Received status code missmatch on atilaClientInsert: " + nextResp.Status)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}

func atilaFetchMessages(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	nextReq, err := http.NewRequest("GET", atilaHost+"/v1/messages?"+request.URL.RawQuery, bytes.NewBuffer(body))
	nextReq.Header.Set("User-Agent", request.UserAgent())
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	client := &http.Client{}

	nextResp, err := client.Do(nextReq)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	defer nextResp.Body.Close()

	respBody, err := ioutil.ReadAll(nextResp.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for the correct status code.
	if nextResp.StatusCode != http.StatusOK && nextResp.StatusCode != http.StatusNoContent {
		fmt.Println("Received status code missmatch on atilaFetchMessages: " + nextResp.Status)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}

func atilaPostMessage(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	nextReq, err := http.NewRequest("POST", atilaHost+"/v1/messages", bytes.NewBuffer(body))
	nextReq.Header.Set("Content-Type", "application/json")
	nextReq.Header.Set("User-Agent", request.UserAgent())
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	client := &http.Client{}

	nextResp, err := client.Do(nextReq)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	defer nextResp.Body.Close()

	respBody, err := ioutil.ReadAll(nextResp.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for the correct status code.
	if nextResp.StatusCode != http.StatusCreated {
		fmt.Println("Received status code missmatch on atilaPostMessage: " + nextResp.Status)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}

func atilaMessageAck(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	nextReq, err := http.NewRequest("POST", atilaHost+"/v1/messages/ack", bytes.NewBuffer(body))
	nextReq.Header.Set("Content-Type", "application/json")
	nextReq.Header.Set("User-Agent", request.UserAgent())
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	client := &http.Client{}

	nextResp, err := client.Do(nextReq)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	defer nextResp.Body.Close()

	respBody, err := ioutil.ReadAll(nextResp.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for the correct status code.
	if nextResp.StatusCode != http.StatusNoContent {
		fmt.Println("Received status code missmatch on atilaMessageAck: " + nextResp.Status)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}
