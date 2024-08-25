package findingsRepo

import (
	"c2/db"
	"c2/models"
)

func GetMultiple() (findings []models.Finding, err error) {
	return findings, db.ORM.Find(&findings).Error
}

func GetMultipleByTag(tag string) (findings []models.Finding, err error) {
	return findings, db.ORM.Where("tag = ?", tag).Find(&findings).Error
}

func Insert(finding *models.Finding) (err error) {
	return db.ORM.Create(&finding).Error
}
