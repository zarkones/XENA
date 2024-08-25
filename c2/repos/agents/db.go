package agentsRepo

import (
	"c2/db"
	"c2/models"
)

func Get(agentID string) (agent models.Agent, err error) {
	return agent, db.ORM.Where("id = ?", agentID).Find(&agent).Error
}

func GetMultiple() (agents []models.Agent, err error) {
	return agents, db.ORM.Find(&agents).Error
}

func Insert(agent *models.Agent) (err error) {
	return db.ORM.Create(&agent).Error
}
