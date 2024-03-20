package handlers

import (
	"errors"

	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/helpers"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var requestBody models.User

	c.BodyParser(&requestBody)

	// Hashing Password
	hashedPassword := helpers.HashPassword([]byte(requestBody.Password))
	if hashedPassword == "" {
		return c.JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"data": errors.New("invalid password").Error(),
		})
	}
	requestBody.Password = hashedPassword

	err := database.DB.Create(&requestBody).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"data": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": requestBody,
	})
}

func Login(c *fiber.Ctx) error {
	var requestBody models.User
	c.BodyParser(&requestBody)

	var user models.User
	err := database.DB.Where("username = ?", requestBody.Username).First(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.StatusNotFound,
			"data": err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(user.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.JSON(fiber.Map{
			"code": fiber.StatusNotFound,
			"data": errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Username)

	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"access_token": access_token,
		"data": user,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}