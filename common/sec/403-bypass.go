package sec

import (
	"common/sec/wordlists"
	"common/slices"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Bypass403(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	errs := []error{}

	c := http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	method := http.MethodGet

	finalURL := url
	if !strings.HasPrefix(finalURL, "http://") && !strings.HasPrefix(finalURL, "https://") {
		finalURL = "https://" + finalURL
	}

	for _, bypassHeader := range wordlists.WebBypass403Headers {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		req.Header.Set(bypassHeader, "127.0.0.1")

		resp, err := c.Do(req)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if resp.Header.Get("Server") != "" {
			analysis = append(analysis, "METHOD: "+method+", URL: "+finalURL+", Server Header: "+resp.Header.Get("Server"))
		}

		if resp.StatusCode == http.StatusForbidden {
			continue
		}

		paths = append(paths, url)

		// Analysis of the response.
		rawRespBody, err := io.ReadAll(resp.Body)
		if err == nil {
			analysis = append(analysis, "METHOD: "+method+", STATUS: "+resp.Status+", LEN: "+strconv.Itoa(len(rawRespBody))+", URL: "+finalURL)

			respBody := string(rawRespBody)
			secrets, _ := FindSecret(&respBody)
			for _, secret := range secrets {
				analysis = append(analysis, "["+method+" - "+finalURL+"]: "+secret)
			}
		}
	}

	paths = slices.Deduplicate(paths)
	analysis = slices.Deduplicate(analysis)

	return paths, analysis, errors.Join(errs...)
}
