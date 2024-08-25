package sec

import (
	"testing"
	"time"
)

// Manually compare time it took to run the test. Stupid solutin for now.
// TODO: Write not a moronic test.
func TestSleep(t *testing.T) {
	start := time.Now().Unix()
	RateLimitedAction(1, func() {
		time.Sleep(time.Second * 2)
	})
	end := time.Now().Unix()
	if end-start > 3 {
		t.Log("unexpected sleep duration")
		t.FailNow()
	}
}
