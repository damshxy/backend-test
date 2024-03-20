package middlewares

import (
	"fmt"

	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/helpers"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
    access_token := c.Get("access_token")

    if len(access_token) == 0 {
        return c.Status(401).SendString("Invalid token: Access token missing")
    }

    checkToken, err := helpers.VerifyToken(access_token)
    if err != nil {
        return c.Status(401).SendString("Invalid token: Failed to verify token")
    }

    fmt.Println(checkToken, "CEKKKK", checkToken["username"])

    var user models.User
    err = database.DB.Where("username = ?", checkToken["username"]).First(&user).Error
    if err != nil {
        fmt.Println(err, "Error fetching user from database")
        return c.Status(401).SendString("Invalid token: Failed to fetch user from database")
    }

    // Set user data in context for future use
    c.Locals("user", user)

    // Continue processing if user is found
    return c.Next()
}
