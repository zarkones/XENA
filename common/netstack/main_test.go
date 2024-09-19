package netstack

import (
	"net"
	"testing"
	"time"
)

func TestSendTcp(t *testing.T) {
	host := "duckduckgo.com"
	req := "GET / HTTP/1.1\r\nHost: " + host + "\r\nConnection: close\r\n\r\n"
	timeout := time.Second * 10

	resp, err := Send(net.JoinHostPort(host, "80"), req, false, false, timeout)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	resp, err = Send(net.JoinHostPort(host, "443"), req, true, false, timeout)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(string(resp))
}
