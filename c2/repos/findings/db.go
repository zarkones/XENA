package findingsRepo

import (
	"c2/db"
	"c2/models"
)

func GetMultiple(offset, limit int) (findings []models.Finding, err error) {
	return findings, db.ORM.Offset(offset).Limit(limit).Find(&findings).Error
}

func GetMultipleByTag(tag string) (findings []models.Finding, err error) {
	return findings, db.ORM.Where("tag = ?", tag).Find(&findings).Error
}

func Insert(finding *models.Finding) (err error) {
	return db.ORM.Create(&finding).Error
}
