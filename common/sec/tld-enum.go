package sec

import (
	"common/slices"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func TldEnum(name string, wordlist *[]string, timeout time.Duration, rpm int) (reachableDomains, analysis []string, err error) {
	reachableDomains = []string{}

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

	for _, tld := range *wordlist {
		RateLimitedAction(sleepAmount, func() {
			finalURL := "https://" + name + tld

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

			reachableDomains = append(reachableDomains, name+tld)
		})
	}

	analysis = slices.Deduplicate(analysis)

	return reachableDomains, analysis, nil
}
