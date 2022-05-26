package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/models"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	config.Db.Where("user_id = ?", id).Find(&links)

	return c.JSON(links)

}
