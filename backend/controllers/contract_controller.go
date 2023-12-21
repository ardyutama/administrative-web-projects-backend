package controllers

import (
	"awp/database"
	"awp/handlers"
	"awp/models"

	"github.com/gofiber/fiber/v2"
)

func AddContracts(c *fiber.Ctx) error {
	u := new(models.Contract)
	handlers.BodyParser(c, &u)
	return handlers.AddEntity(c, &u)
}

func GetContracts(c *fiber.Ctx) error {
	u := []models.Contract{}
	database.DB.Preload("Service").Find(&u)
	return c.JSON(u)
}
