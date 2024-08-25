package sec

import (
	"fmt"
	"testing"
	"time"
)

func TestSubdomainsEnum(t *testing.T) {
	wordlist := []string{"www", "doesnotexist", "ftp"}
	target := "duckduckgo.com"

	domains, _, err := SubdomainEnum(target, &wordlist, time.Second*5, 1000)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(domains) == 0 {
		t.FailNow()
	}

	fmt.Println(domains)
}
