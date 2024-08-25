package sec

import (
	"common/sec/wordlists"
	"common/slices"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

var redirectHeaders = []string{
	"X-Forwarded-For-Original",
	"X-Forwarded-For",
	"X-Forwarded-Server",
	"X-Forwarded",
	"X-Forwarder-For",
	"X-Host",
}

func HostHeaderInjection(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
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

	finalUrl := url
	if !strings.HasPrefix(finalUrl, "http://") && !strings.HasPrefix(finalUrl, "https://") {
		finalUrl = "https://" + finalUrl
	}

	fictionalDomain := slices.Rand(&wordlists.WordsTop850) + slices.Rand(&wordlists.WordsTop850) + slices.Rand(&wordlists.TLDsCommon)

	func() {
		method := http.MethodGet

		req, err := http.NewRequest(method, finalUrl, nil)
		if err != nil {
			errs = append(errs, err)
			return
		}

		req.Host = fictionalDomain

		resp, err := c.Do(req)
		if err != nil {
			errs = append(errs, err)
			return
		}

		if !strings.Contains(resp.Header.Get("Location"), fictionalDomain) {
			return
		}

		paths = append(paths, finalUrl)
		analysis = append(analysis, "METHOD: "+method+", URL: "+finalUrl+" Redirected with status: "+resp.Status)
	}()

	func() {
		method := http.MethodGet

		for _, redHeader := range redirectHeaders {
			req, err := http.NewRequest(method, finalUrl, nil)
			if err != nil {
				errs = append(errs, err)
				return
			}

			req.Header.Set(redHeader, fictionalDomain)

			resp, err := c.Do(req)
			if err != nil {
				errs = append(errs, err)
				return
			}

			rawRespBody, err := io.ReadAll(resp.Body)
			if err != nil {
				errs = append(errs, err)
				return
			}

			if !strings.Contains(string(rawRespBody), fictionalDomain) {
				return
			}

			paths = append(paths, finalUrl)
			analysis = append(analysis, "METHOD: "+method+", URL: "+finalUrl+" Included fictional domain '"+resp.Status+"' in the response body")
		}
	}()

	paths = slices.Deduplicate(paths)

	return paths, analysis, errors.Join(errs...)
}
