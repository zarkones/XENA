package sec

import (
	"testing"
	"time"
)

func TestHttpScanner(t *testing.T) {
	requests := []string{
		"GET /v1/users?limit=20&offset=0&status=all HTTP/1.1\r\nHost: localhost:8080\r\nConnection: close\r\n\r\n",
		"POST /v1/users HTTP/1.1\r\nHost: localhost:8080\r\nContent-Type: application/json\r\nConnection: close\r\n\r\n{\"foo\":\"bar\",\"foo2\":\"bar2222\",\"abc\":123,\"xyz\":false,\"nono\":{\"foo\":\"bar\"}}",
	}

	for _, req := range requests {
		_, err := HttpScanner(&req, false, time.Second, 50)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}
}
