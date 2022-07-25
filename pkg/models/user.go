package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email" gorm:"unique"`
	Password     []byte   `json:"-"`
	IsAmbassador bool     `json:"-"`
	Revenue      *float64 `json:"revenue,omitempty" gorm:"-"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(User{})
// }

func (user *User) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashPassword

}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

type Admin User

func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   admin.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}

	admin.Revenue = &revenue

}

type Ambassador User

func (ambassador *Ambassador) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   ambassador.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AmbassadorRevenue
		}
	}

	ambassador.Revenue = &revenue

}
