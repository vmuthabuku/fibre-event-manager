package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	Id           uint
	FirstName    string
	LastName     string
	Email        string
	Password     string
	isAmbassador string
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(User{})
// }
