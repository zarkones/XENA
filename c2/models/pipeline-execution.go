package models

import (
	"strings"
	"time"
)

type PipelineRun struct {
	ID               string `json:"id" gorm:"primaryKey"`
	PipelineID       string `json:"pipelineId" gorm:"primaryKey"`
	FinishedPipeline string `json:"finishedPipeline"`
	FinishedAt       int64  `json:"finishedAt"`
	FinishedAtLabel  string `json:"finishedAtLabel"`
	ExecutedAt       int64
}

func (p *PipelineRun) Finished() {
	now := time.Now()
	p.FinishedAt = now.Unix()
	p.FinishedAtLabel = strings.Split(now.String(), ".")[0]
}
