package models

import "time"

type ProxyReq struct {
	ID              int64 `gorm:"priamryKey"`
	SessionID       int64
	Status          int
	Host            string
	Method          string
	Path            string
	Query           string
	Body            bool
	ReqContentType  string
	RespContentType string
	ReqLength       int
	RespLength      int
	RawReq          string
	RawResp         string
	Time            time.Time

	// Analysis flags:
	Analyzed bool
}
