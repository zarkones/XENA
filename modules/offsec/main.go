package offsec

import (
	"common/debug"
	"encoding/json"
	"strings"

	"github.com/gocolly/colly"
	c2api "github.com/zarkones/xena-client"
)

const PLUGIN_NAME = "PIPELINES"

func main() {}

//export Init
func Init(input string) (output string) {
	var pipeline c2api.Pipeline

	if err := json.Unmarshal([]byte(input), &pipeline); err == nil {
		debug.Println(PLUGIN_NAME, "interpreting command:", input)

		run := runPipeline(pipeline)

		runJson, err := json.Marshal(&run)
		if err != nil {
			return err.Error()
		}

		return string(runJson)
	}

	return ""
}

func webCrawl(domain string) (output string) {
	c := colly.NewCollector()

	foundUrls := []string{}

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("href") == "" {
			return
		}
		normalizedTarget := strings.ReplaceAll(domain, "http://", "")
		normalizedTarget = strings.ReplaceAll(normalizedTarget, "https://", "")
		if !strings.Contains(e.Attr("href"), normalizedTarget) {
			foundUrls = append(foundUrls, "Skipping: "+e.Attr("href"))
			return
		}
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		foundUrls = append(foundUrls, r.URL.String())
	})

	if err := c.Visit(domain); err != nil {
		return err.Error()
	}

	return strings.Join(foundUrls, "\n")
}
