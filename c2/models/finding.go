package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Finding struct {
	ID         string `json:"id"`
	PipelineID string `json:"pipelineId"`
	Tag        string `json:"tag"`
	Data       string `json:"data"`
}

func (m *Finding) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}
