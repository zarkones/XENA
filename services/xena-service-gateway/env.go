package main

import "os"

// Gateway configuration.

// Port on which the server runs on.
var port string = "60606"

var atilaHost string = os.Getenv("ATILA_HOST")
var domenaHost string = os.Getenv("DOMENA_HOST")
