package attacksRepo

import (
	"c2/db"
	"c2/models"
)

func Get(attackID string) (attack models.Attack, err error) {
	return attack, db.ORM.Where("id = ?", attackID).First(&attack).Error
}

func Delete(attackID string) (err error) {
	return db.ORM.Where("id = ?", attackID).Delete(&models.Attack{}).Error
}

func GetMultiple() (attacks []models.Attack, err error) {
	return attacks, db.ORM.Find(&attacks).Error
}

func Upsert(attack *models.Attack) (err error) {
	return db.ORM.Save(&attack).Error
}
