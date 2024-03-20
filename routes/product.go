package routes

import (
	"github.com/damshxy/manajemen-backend-test/handlers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(r fiber.Router) {
	r.Get("/products", handlers.GetProducts)
	r.Get("/product/:id", handlers.GetProductByID)
	r.Post("/product", handlers.CreateProduct)
	r.Patch("/product/:id", handlers.UpdateProduct)
	r.Delete("/product/:id", handlers.DeleteProduct)
}