package sec

import (
	"net/url"
	"strings"
	"time"
)

// TODO: WebUrlFuzzer
func WebUrlFuzzer(method, targetURL string, timeout time.Duration, rpm int) (urls, analysis []string, err error) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, nil, err
	}

	if parsedURL.RawQuery == "" {
		return []string{}, []string{}, nil
	}

	keyValuePairs := strings.Split(parsedURL.RawQuery, "&")

	params := map[string]string{}

	for _, keyValuePair := range keyValuePairs {
		deconstructed := strings.Split(keyValuePair, "=")
		key := ""
		val := ""
		for partIndex, part := range deconstructed {
			if partIndex == 0 {
				key = part
				continue
			}
			val += part
		}
		params[key] = val
	}

	// targetUrlWithoutParams := strings.Split(targetURL, "?")[0]

	// TODO: Read about backslash scanner and implement it here.
	// https://portswigger.net/research/backslash-powered-scanning-hunting-unknown-vulnerability-classes

	return nil, nil, nil
}
