package state

import (
	xenaC2 "github.com/zarkones/xena-client"
)

var Files = []xenaC2.File{}
var Pipelines = []xenaC2.Pipeline{}
var Agents = []xenaC2.Agent{}
var AuthToken = ""
var C2Host = "http://127.0.0.1:8080"
