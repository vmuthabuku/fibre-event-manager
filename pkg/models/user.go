package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName    string
	LastName     string
	Email        string
	Password     []byte
	IsAmbassador bool
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
