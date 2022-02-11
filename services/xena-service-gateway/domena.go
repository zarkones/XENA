package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func domenaServiceInsert(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	nextReq, err := http.NewRequest("POST", domenaHost+"/v1/services", bytes.NewBuffer(body))
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
		fmt.Println("Received status code missmatch on domenaServiceInsert: " + nextResp.Status)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}
