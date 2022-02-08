package main

import (
	"fmt"
	"net/http"
)

func main() {
	initRoutes()

	// Listen on the port.
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
