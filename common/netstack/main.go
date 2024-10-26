package netstack

import (
	"bufio"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func Send(host, request string, secure, allowInsecure bool, timeout time.Duration) (rawResp []byte, err error) {
	scheme := "https"
	if !secure {
		scheme = "http"
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: allowInsecure,
			},
		},
		Timeout: timeout,
	}

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(request)))
	if err != nil {
		return nil, err
	}

	req.RequestURI, req.URL.Scheme, req.URL.Host = "", scheme, req.Host

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return httputil.DumpResponse(resp, true)
}
