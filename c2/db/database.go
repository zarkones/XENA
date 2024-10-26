package db

import (
	"c2/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var ORM *gorm.DB

func Init(dbName string) error {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&models.Agent{},
		&models.Message{},
		&models.Pipeline{},
		&models.Finding{},
		&models.Target{},
		&models.Attack{},
		&models.PipelineRun{},
		&models.File{},
		&models.ProxyReq{},
	); err != nil {
		return err
	}

	ORM = db

	return nil
}
