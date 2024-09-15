package proxyRepo

import (
	"c2/db"
	"c2/models"
	"strings"
)

func GetMultiple(offset, limit int, orderBy, orderDirection string) (requests []models.ProxyReq, err error) {
	tx := db.ORM.
		Offset(offset).
		Limit(limit)

	if len(orderBy) != 0 && len(orderDirection) != 0 {
		if orderDirection[0] == 'A' || orderDirection[0] == 'a' {
			orderDirection = "ASC"
		} else {
			orderDirection = "DESC"
		}
		q := strings.ToLower(orderBy) + " " + orderDirection
		tx.Order(q)
	}

	return requests, tx.Find(&requests).Error
}

func Insert(req *models.ProxyReq) (err error) {
	return db.ORM.Create(req).Error
}

func UpdateRawResp(sessionID int64, statusCode int, rawResp *string) (err error) {
	var req models.ProxyReq
	if err := db.ORM.Where("session_id = ?", sessionID).First(&req).Error; err != nil {
		return err
	}

	req.Status = statusCode
	req.RawResp = *rawResp
	req.RespLength = len(*rawResp)

	return db.ORM.Save(&req).Error
}
