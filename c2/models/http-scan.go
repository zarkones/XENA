package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HttpScan is meant to provide to C2 the operator's intent to scan a http request.
// C2 should based on the internal scanner engine generate HttpScanTask(s), and
// based on their responses potentially generate more of them. Inspired by the
// backslash powered scanner by Mr. James Kettle. It should not spam endpoints
// with 1 million payloads, rather it should try to find suspicious behavior.
// As this is behavior scanner, NOT a vulnerability scanner.
type HttpScan struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ReqID     int64     `json:"reqId"`
	AgentIDs  string    `json:"agentIds"`
	CreatedAt time.Time `json:"createdAt"`
}

func (m *HttpScan) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now()
	return nil
}

// HttpScanTask has a simple purpose. Here is a http request with some payload,
// issue the request and send back the response. At least that's what happens
// from an agent's perspective.
type HttpScanTask struct {
	ScanID      string `json:"scanId" gorm:"primaryKey"`
	AgentID     string `json:"agentId" gorm:"primaryKey"`
	ReqID       int64  `json:"reqId" gorm:"primaryKey"`
	Payload     string `json:"payload"`
	RawRequest  string `json:"rawRequest"`
	RawResponse string `json:"rawResponse"`
}
