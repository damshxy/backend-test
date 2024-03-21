package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Get("/users", handlers.GetAllUser)
}