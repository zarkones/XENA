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
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	nextReq, err := http.NewRequest("POST", domenaHost+"/v1/services", bytes.NewBuffer(body))
	nextReq.Header.Set("Content-Type", "application/json")
	nextReq.Header.Set("User-Agent", request.UserAgent())
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	nextResp, err := client.Do(nextReq)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer nextResp.Body.Close()

	respBody, err := ioutil.ReadAll(nextResp.Body)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(nextResp.StatusCode)
	response.Write(respBody)
}
