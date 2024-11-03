package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Finding struct {
	ID         string `json:"id" gorm:"primaryKey"`
	PipelineID string `json:"pipelineId"`
	RequestID  int64  `json:"requestId"`
	Tag        string `json:"tag"`
	Data       string `json:"data"`
}

func (m *Finding) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}
