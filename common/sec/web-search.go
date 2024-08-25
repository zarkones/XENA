package sec

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const DuckDuckGoURL = "https://html.duckduckgo.com/html/"

var regexDuckGo = func() *regexp.Regexp {
	exp, _ := regexp.Compile(`<a class="result__url" href="(.*?)>`)
	return exp
}()

func WebSearch(query string, page int, timeout time.Duration, rpm int) (urls, analysis []string, err error) {
	c := http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	analysis = []string{}

	serializedPage := func() string {
		switch page {
		case 0, 1:
			return ""
		case 2:
			return "29"
		default:
			return strconv.Itoa(page*50 + 29)
		}
	}()

	req, err := http.NewRequest(http.MethodPost, DuckDuckGoURL, bytes.NewReader([]byte("b=&s="+serializedPage+"&q="+query)))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://html.duckduckgo.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Origin", "https://html.duckduckgo.com")
	req.Header.Set("Cookie", "kl=wt-wt")

	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}

	rawRespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	foundURLs := regexDuckGo.FindAllString(string(rawRespBody), -1)
	urls = make([]string, len(foundURLs))
	analysis = make([]string, len(foundURLs))

	for i, foundURL := range foundURLs {
		foundURL = strings.TrimPrefix(foundURL, `<a class="result__url" href="`)
		foundURL = strings.TrimSuffix(foundURL, `">`)
		urls[i] = foundURL
		analysis[i] = foundURL
	}

	return urls, analysis, nil
}
