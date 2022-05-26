package main

import (
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
	"github.com/vmuthabuku/fibre-event-manager/pkg/utils"
)

func main() {
	config.Connect()
	for i := 0; i < 10; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		products := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		ambassador.SetPassword("1234")

		config.Db.Create(&ambassador)
		config.Db.Create(&products)

	}

	utils.PopulateOrders()
	fmt.Println("Created Successfully")
}
