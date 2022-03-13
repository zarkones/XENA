package services

import (
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

// HTTP contract of DuckDuckGos search endpoint.
type DuckDuckGoSearch struct {
	Q string `json:"q"` // Search term.
}

func DuckitAndSleep(term, domain, inurl string, minSleepSec, maxSleepSec int) ([]string, error) {
	query := term

	if len(inurl) != 0 {
		query += " inurl:" + inurl
	}
	if len(domain) != 0 {
		query += " site:" + domain
	}

	results, err := Duckit(query)
	if err != nil {
		return results, nil
	}

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Second * time.Duration(rand.Intn(maxSleepSec-minSleepSec)+maxSleepSec))

	return results, nil
}

// duckit performs a web search using the DuckDuckGo.
// Returns a slice of strings representing urls.
func Duckit(term string) ([]string, error) {
	var searchResults []string

	request, err := http.NewRequest("POST", "https://html.duckduckgo.com/html?q="+term, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0")
	request.Header.Set("Origin", "https://html.duckduckgo.com")
	request.Header.Set("Connection", "close")
	if err != nil {
		return searchResults, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return searchResults, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return searchResults, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return searchResults, err
	}

	r := regexp.MustCompile(`result__url" href="(.*?)">`)
	matches := r.FindAllString(string(bodyBytes), -1)

	for _, url := range matches {
		searchResults = append(searchResults, url[19:len(url)-2])
	}

	return searchResults, nil
}
