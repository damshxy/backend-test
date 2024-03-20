package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(r fiber.Router) {
	r.Get("/orders", handlers.GetOrders)
	r.Get("/order/:id", handlers.GetOrderByID)
	r.Post("/order", handlers.CreateOrder)
	r.Patch("/order/:id", handlers.UpdateOrder)
	r.Delete("/order/:id", handlers.DeleteOrder)
}