package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(r fiber.Router) {
	r.Get("/orders", handlers.GetOrders)
	r.Patch("/order/:id/process", handlers.ProcessOrder)
	r.Patch("/order/:id/complete", handlers.CompleteOrder)
}