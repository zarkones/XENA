package models

import "time"

type ProxyReq struct {
	ID        int64 `gorm:"priamryKey"`
	SessionID int64
	Host      string
	Method    string
	Path      string
	Query     string
	Body      bool
	Length    int
	Raw       string
	Time      time.Time
}
