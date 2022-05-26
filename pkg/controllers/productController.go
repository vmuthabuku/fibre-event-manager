package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	config.Db.Find(&products)

	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	config.Db.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)

	config.Db.Find(&product)

	return c.JSON(product)
}
