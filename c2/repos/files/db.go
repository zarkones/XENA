package filesRepo

import (
	"c2/db"
	"c2/models"
	"time"
)

func Get(id string) (file models.File, err error) {
	return file, db.ORM.Where("id = ?", id).First(&file).Error
}

func GetMultiple() (files []models.File, err error) {
	return files, db.ORM.Find(&files).Error
}

func Delete(id string) (err error) {
	return db.ORM.Where("id = ?", id).Delete(&models.File{}).Error
}

func Insert(file *models.File) (err error) {
	return db.ORM.Create(&file).Error
}

func SetUploaded(id string) (err error) {
	file, err := Get(id)
	if err != nil {
		return err
	}
	if file.Uploaded {
		return nil
	}
	file.Uploaded = true
	file.UploadedAt = time.Now()
	return db.ORM.Save(&file).Error
}
