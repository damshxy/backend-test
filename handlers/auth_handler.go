package handlers

import (
	"errors"

	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/helpers"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	// Save user to database
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

	// create a new validator
	validate := validator.New()

	// Validate the user struct
	if err := validate.Struct(requestBody); err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"data": err.Error(),
		})
	}

	// Check Password
	checkPassword := helpers.ComparePassword([]byte(user.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.JSON(fiber.Map{
			"code": fiber.StatusNotFound,
			"data": errors.New("invalid credentials").Error(),
		})
	}


	access_token := helpers.SignToken(requestBody.Username)

	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"access_token": access_token,
		"data": user,
	})
}

func Profile(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(jwt.MapClaims)
	
	username := userClaims["username"].(string)

	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

    return c.JSON(fiber.Map{
        "message": "User profile fetched successfully",
        "data":    user,
    })
}



func ResetPassword(c *fiber.Ctx) error {
	var resetPassword models.ResetPassword
	if err := c.BodyParser(&resetPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword := helpers.HashPassword([]byte(resetPassword.NewPassword))

	if err := database.DB.Model(&models.User{}).Where("username = ?", resetPassword.Username).Update("password", hashedPassword).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Password reset successfully",
	})
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(jwt.MapClaims)

	c.ClearCookie("access_token")

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
		"user": user,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}