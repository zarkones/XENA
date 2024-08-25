package sec

import (
	"fmt"
	"testing"
	"time"
)

func TestTldEnum(t *testing.T) {
	wordlist := []string{".com", ".doesnotexist", ".spaceship123"}
	target := "duckduckgo"

	domains, _, err := TldEnum(target, &wordlist, time.Second*5, 5)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(domains) == 0 {
		t.FailNow()
	}

	fmt.Println(domains)
}
