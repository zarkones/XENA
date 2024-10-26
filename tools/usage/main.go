package usage

import (
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Details   string `json:"details"`
	Err       string `json:"err"`
	Timestamp int64  `json:"timestamp"`
}

func NewEvent(name, details string, maybeErr error) error {
	logsPath := os.Getenv("XENA_LOGS_PATH")
	if logsPath == "" {
		logsPath = "xena.logs"
	}
	f, err := os.OpenFile(logsPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	e := ""
	if maybeErr != nil {
		e = maybeErr.Error()
	}
	jsonEvent, err := json.Marshal(Event{
		ID:        uuid.NewString(),
		Name:      name,
		Details:   details,
		Err:       e,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	if _, err := f.WriteString(string(jsonEvent) + "\n"); err != nil {
		return err
	}
	return nil
}
