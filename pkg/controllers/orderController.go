package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

func Order(c *fiber.Ctx) error {

	var orders []models.Order

	config.Db.Find(&orders)

	return c.JSON(orders)

}
