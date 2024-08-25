package models

import "time"

type File struct {
	ID                string    `json:"id" gorm:"primaryKey"`
	UploadedByAgentID string    `json:"uploadedByAgentId"`
	OriginalName      string    `json:"originalName"`
	StorageName       string    `json:"storageName"`
	Uploaded          bool      `json:"uploaded"`
	UploadedAt        time.Time `json:"uploadedAt"`
}
