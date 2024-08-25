package data

import "net/url"

var XenaMobileURL = func() *url.URL {
	xenaMobileURL, _ := url.Parse("https://zarkones.itch.io/xena-mobile")
	return xenaMobileURL
}()
