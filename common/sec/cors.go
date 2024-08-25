package sec

import (
	"common/debug"
	"common/sec/wordlists"
	"common/slices"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	urlUtil "net/url"
	"strconv"
	"strings"
	"time"
)

func AnalyzeCors(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	c := http.Client{
		Timeout: timeout,
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// return http.ErrUseLastResponse
		// },
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url = strings.TrimSuffix(url, "/")
	analysis = []string{}
	errs := []error{}

	sleepAmount := GetSleep(rpm)

	parsedURL, err := urlUtil.Parse(url)
	if err != nil {
		return nil, nil, err
	}
	targetHost := parsedURL.Hostname()
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		targetHost = url
	}

	fictionalAtkDomain := slices.Rand(&wordlists.WordsTop850) + slices.Rand(&wordlists.WordsTop850) + slices.Rand(&wordlists.TLDsCommon)

	hosts := []string{
		fictionalAtkDomain,
		targetHost + "." + fictionalAtkDomain,
		targetHost + fictionalAtkDomain,
		slices.Rand(&wordlists.WordsTop850) + slices.Rand(&wordlists.WordsTop850) + "." + targetHost,
		"null",
	}

	method := http.MethodGet

	for _, host := range hosts {
		host = "https://" + host

		finalURL := url
		if !strings.HasPrefix(url, "http") {
			finalURL = "https://" + finalURL
		}

		RateLimitedAction(sleepAmount, func() {
			debug.Println("CORS:", method, finalURL, "PAYLOAD:", host)

			req, err := http.NewRequest(method, finalURL, nil)
			if err != nil {
				debug.Println("sec.AnalyzeCors:", method, finalURL, err)
				errs = append(errs, err)
				return
			}

			req.Header.Set("Origin", host)

			resp, err := c.Do(req)
			if err != nil {
				debug.Println("sec.AnalyzeCors:", method, finalURL, err)
				errs = append(errs, err)
				return
			}

			if resp.Header.Get("Server") != "" {
				analysis = append(analysis, "METHOD: "+method+", URL: "+finalURL+", Server Header: "+resp.Header.Get("Server"))
			}

			acao := resp.Header.Get("Access-Control-Allow-Origin")
			acac := resp.Header.Get("Access-Control-Allow-Credentials")

			// Analysis of the response.
			rawRespBody, err := io.ReadAll(resp.Body)
			if err == nil {
				analysis = append(analysis, "ORIGIN: "+host+", STATUS: "+resp.Status+", LEN: "+strconv.Itoa(len(rawRespBody))+", HEADERS: "+acao+" & "+acac+", URL: "+finalURL)

				respBody := string(rawRespBody)
				secrets, _ := FindSecret(&respBody)
				for _, secret := range secrets {
					analysis = append(analysis, "["+method+" - "+finalURL+"]: "+secret)
				}
			}

			if acao != host && acao != "*" {
				return
			}

			paths = append(paths, finalURL)
		})
	}

	paths = slices.Deduplicate(paths)
	analysis = slices.Deduplicate(analysis)

	return paths, analysis, errors.Join(errs...)
}
