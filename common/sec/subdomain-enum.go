package sec

import (
	"common/slices"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func SubdomainEnum(domain string, wordlist *[]string, timeout time.Duration, rpm int) (reachableSubdomains, analysis []string, err error) {
	reachableSubdomains = []string{}

	c := http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	sleepAmount := GetSleep(rpm)
	analysis = []string{}

	for _, subdomain := range *wordlist {
		RateLimitedAction(sleepAmount, func() {
			finalURL := "https://" + subdomain + "." + domain

			resp, err := c.Get(finalURL)
			if err != nil {
				return
			}

			if resp.Header.Get("Server") != "" {
				analysis = append(analysis, "URL: "+finalURL+", Server Header: "+resp.Header.Get("Server"))
			}

			// Analysis of the response.
			rawRespBody, err := io.ReadAll(resp.Body)
			if err == nil {
				respBody := string(rawRespBody)
				newSecrets, _ := FindSecret(&respBody)
				for _, secret := range newSecrets {
					analysis = append(analysis, "["+finalURL+"]: "+secret)
				}
			}

			reachableSubdomains = append(reachableSubdomains, subdomain+"."+domain)
		})
	}

	analysis = slices.Deduplicate(analysis)

	return reachableSubdomains, analysis, nil
}
