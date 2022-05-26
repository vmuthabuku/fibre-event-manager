package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

func Ambassador(c *fiber.Ctx) error {
	var users []models.User

	config.Db.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}
