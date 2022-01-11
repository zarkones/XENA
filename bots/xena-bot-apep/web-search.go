package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// HTTP contract of DuckDuckGos search endpoint.
type DuckDuckGoSearch struct {
	Q string `json:"q"` // Search term.
}

// duckit performs a web search using the DuckDuckGo.
// Returns a slice of strings representing urls.
func duckit(term string) ([]string, error) {
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

	fmt.Println(string(bodyBytes))

	r := regexp.MustCompile(`result__url" href="(.*?)">`)
	matches := r.FindAllString(string(bodyBytes), -1)

	for _, url := range matches {
		searchResults = append(searchResults, url[19:len(url)-2])
	}

	return searchResults, nil
}
