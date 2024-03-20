package routes

import "github.com/gofiber/fiber/v2"

func MainRoutes(app fiber.Router) {
	ProductRoutes(app)
	UserRoutes(app)
	OrderRoutes(app)
}