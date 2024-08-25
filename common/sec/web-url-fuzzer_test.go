package sec

import (
	"net/http"
	"testing"
	"time"
)

func TestWebUrlFuzzer(t *testing.T) {
	const url = "https://xena.network?q=123&foo=bar&x="

	WebUrlFuzzer(http.MethodGet, url, time.Second*10, 10)
}
