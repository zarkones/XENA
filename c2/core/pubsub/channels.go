package pubsub

import (
	xenaC2 "github.com/zarkones/xena-client"
)

var Messages = map[string]chan xenaC2.Message{}
