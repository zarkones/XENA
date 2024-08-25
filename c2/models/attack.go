package models

type Attack struct {
	ID       string `json:"id"`
	AgentID  string `json:"agentId" gorm:"primaryKey"`
	TargetID string `json:"targetId" gorm:"primaryKey"`
	Comment  string `json:"comment"`
}
