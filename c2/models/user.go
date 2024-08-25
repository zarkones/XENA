package models

type User struct {
	Name         string `json:"name" gorm:"primaryKey"`
	PasswordHash string `json:"passwordHash"`
}
