package handlers

import (
	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	err := database.DB.Order("id").Find(&orders).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error fetching orders",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Orders fetched successfully",
		"data":    orders,
	})
}

func ProcessOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order

	err := database.DB.First(&order, orderID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Order not found",
			})
		}
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"message": "Failed to process order",
		})
	}

	order.Status = "processed"
	err = database.DB.Save(&order).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"message": "Failed to process order",
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"message": "Order processed successfully",
		"data":    order,
	})
}

func CompleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order

	err := database.DB.First(&order, orderID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Order not found",
			})
		}
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"message": "Failed to complete order",
		})
	}

	order.Status = "Completed"
	err = database.DB.Save(&order).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"message": "Failed to complete order",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order completed successfully",
		"data": order,
	})
}