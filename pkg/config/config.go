package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

var (
	Db *gorm.DB
)

func Connect() {

	var err error
	d, err := gorm.Open("mysql", "root:password@/event_manager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db = d
}

func AutoMigrate() {
	Db.AutoMigrate(models.User{}, models.Product{}, models.Link{}, models.Order{}, models.OrderItem{})
}

// func GetDB() *gorm.DB {
// 	return db
// }
