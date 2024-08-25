package targetsRepo

import (
	"c2/db"
	"c2/models"
)

func GetByValue(value string) (target models.Target, err error) {
	return target, db.ORM.Where("value = ?", value).First(&target).Error
}

func GetDomain(domain string) (target models.Target, err error) {
	return target, db.ORM.Where("value = ?", domain).Where("type = ?", "DOMAIN").First(&target).Error
}

func Get(targetName string) (target models.Target, err error) {
	return target, db.ORM.Where("name = ?", targetName).First(&target).Error
}

func Delete(targetID string) (err error) {
	return db.ORM.Where("id = ?", targetID).Delete(&models.Target{}).Error
}

func GetMultiple() (targets []models.Target, err error) {
	return targets, db.ORM.Find(&targets).Error
}

func GetChildren(targetID string) (targets []models.Target, err error) {
	return targets, db.ORM.Where("parent_id = ?", targetID).Find(&targets).Error
}

func GetMultipleByIDs(targetIDs []string) (targets []models.Target, err error) {
	return targets, db.ORM.Where("parent_id = ?", "").Where(targetIDs).Find(&targets).Error
}

func Upsert(target *models.Target) (err error) {
	return db.ORM.Save(&target).Error
}
