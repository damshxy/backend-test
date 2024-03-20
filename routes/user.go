package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/damshxy/manajemen-backend-test/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	userGroup := r.Group("/user")

	userGroup.Post("/register", handlers.Register)
	userGroup.Post("/login", handlers.Login)
	userGroup.Get("/auth", middlewares.Authentication, handlers.Auth)
}