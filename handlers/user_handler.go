package handlers

import (
	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User

	err := database.DB.Order("id").Find(&users).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    users,
	})
}