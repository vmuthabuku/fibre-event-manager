package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

func main() {
	config.Connect()
	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("1234")

		config.Db.Create(&ambassador)

	}
	fmt.Println("Created Successfully")
}
