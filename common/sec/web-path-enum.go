package sec

import (
	"common/debug"
	"common/sec/wordlists"
	"common/slices"
	"crypto/tls"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func WebFindGitDir(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathGit, timeout, rpm)
}

func WebFindAdminPanels(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathAdmin, timeout, rpm)
}

func WebFindLogs(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathLogs, timeout, rpm)
}

func WebFindBackups(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathBackups, timeout, rpm)
}

func WebFindDevLeftover(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathDevLeftover, timeout, rpm)
}

func WebFindTop10K(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	return WebPathEnum(http.MethodGet, url, &wordlists.WebPathTop10k, timeout, rpm)
}

func WebFindFileUpload(url string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
	newPaths, newAnalysis, _ := WebPathEnum(http.MethodPost, url, &wordlists.WebPathFileUpload, timeout, rpm)
	analysis = newAnalysis
	paths = newPaths

	newPaths, newAnalysis, _ = WebPathEnum(http.MethodPut, url, &wordlists.WebPathFileUpload, timeout, rpm)
	analysis = append(analysis, newAnalysis...)
	paths = append(paths, newPaths...)

	newPaths, newAnalysis, _ = WebPathEnum(http.MethodPatch, url, &wordlists.WebPathFileUpload, timeout, rpm)
	analysis = append(analysis, newAnalysis...)
	paths = append(paths, newPaths...)

	paths = slices.Deduplicate(paths)

	return paths, analysis, err
}

func WebPathEnum(method, url string, wordlist *[]string, timeout time.Duration, rpm int) (paths, analysis []string, err error) {
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

	url = strings.TrimSuffix(url, "/")
	analysis = []string{}

	sleepAmount := GetSleep(rpm)

	for _, path := range *wordlist {
		path = strings.TrimPrefix(path, "/")
		finalURL := url + "/" + path
		if !strings.HasPrefix(url, "http") {
			finalURL = "https://" + finalURL
		}

		RateLimitedAction(sleepAmount, func() {
			debug.Println("Probbing:", method, finalURL)

			req, err := http.NewRequest(method, finalURL, nil)
			if err != nil {
				debug.Println("sec.WebPathEnum:", method, finalURL, err)
				return
			}

			resp, err := c.Do(req)
			if err != nil {
				debug.Println("sec.WebPathEnum:", method, finalURL, err)
				return
			}
			if resp.StatusCode == http.StatusNotFound {
				return
			}

			if resp.Header.Get("Server") != "" {
				analysis = append(analysis, "METHOD: "+method+", URL: "+finalURL+", Server Header: "+resp.Header.Get("Server"))
			}

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

			paths = append(paths, finalURL)
		})
	}

	return paths, analysis, nil
}
