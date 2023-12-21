package controllers

import (
	"awp/database"
	"awp/handlers"
	"awp/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetVMStatuses(c *fiber.Ctx) error {
	u := []models.VMStatus{}
	database.DB.Find(&u)
	return c.JSON(u)
}

func AddVMStatuses(c *fiber.Ctx) error {
	u := new(models.VMStatus)
	handlers.BodyParser(c, &u)
	return handlers.AddUniqueEntity(c, &u)
}

func DeleteVMStatuses(c *fiber.Ctx) error {
	u := new(models.VMStatus)

	id := c.Params("id")

	err := database.DB.First(&u, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if err := database.DB.Delete(&u).Error; err != nil {
		// Handle the error during deletion
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete record"})
	}

	return c.Status(200).JSON("deleted")
}

func EditVMStatuses(c *fiber.Ctx) error {
	u := new(models.VMStatus)

	handlers.BodyParser(c, &u)

	id := c.Params("id")

	exist := new(models.VMStatus)
	res := database.DB.Model(&models.VMStatus{}).Where("id = ?", id).First(exist)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// Handle the case where the record doesn't exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res = database.DB.Model(&models.VMStatus{}).Where("id = ?", id).Update("name", u.Name)
	if res.Error != nil {
		// Handle the error during update
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update record"})
	}

	return c.Status(400).JSON("updated")
}
