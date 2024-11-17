package httpScansRepo

import (
	"c2/db"
	"c2/models"
)

func GetMultiple(limit, offset int) (scans []models.HttpScan, err error) {
	return scans, db.ORM.Offset(offset).Limit(limit).Find(&scans).Error
}

func Insert(requestID int64, agentIDs string) (err error) {
	return db.ORM.Create(&models.HttpScan{
		ReqID:    requestID,
		AgentIDs: agentIDs,
	}).Error
}
