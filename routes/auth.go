package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/damshxy/manajemen-backend-test/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	r.Post("/logout", middlewares.Authentication, handlers.Logout)
	r.Patch("/resetPassword", handlers.ResetPassword)
	r.Get("/profile", middlewares.Authentication, handlers.Profile)
	r.Get("/auth", middlewares.Authentication, handlers.Auth)
}