package proxyRepo

import (
	"c2/db"
	"c2/models"
)

func GetMultiple(offset, limit int) (requests []models.ProxyReq, err error) {
	return requests, db.ORM.
		Offset(offset).
		Limit(limit).
		Find(&requests).Error
}

func Insert(req *models.ProxyReq) (err error) {
	return db.ORM.Create(req).Error
}
