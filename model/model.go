package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey"`
	Password string
	Cliptext string
	Expired int `gorm:"default:30"`
}
