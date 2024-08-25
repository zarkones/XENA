package sec

import "time"

func GetSleep(rpm int) (sleepAmount float32) {
	return 60 / float32(rpm)
}

func RateLimitedAction(sleepAmount float32, action func()) {
	// We don't want to sleep for less than 100ms.
	if sleepAmount < 0.1 {
		action()
		return
	}

	start := time.Now()
	action()
	end := time.Now()
	cooldown := start.Add(time.Second * time.Duration(sleepAmount))
	if end.After(cooldown) {
		return
	}
	time.Sleep(time.Second * time.Duration(cooldown.Unix()-end.Unix()))
}
