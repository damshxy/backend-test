package handlers

import (
	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
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

func GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order

	err := database.DB.First(&order, orderID).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error fetching order",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Order fetched successfully",
		"data":    order,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	var requestBody models.Order
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing request body",
		})
	}

	err = database.DB.Create(&requestBody).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error creating order",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Order created successfully",
		"data":    requestBody,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order

	err := c.BodyParser(&order)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing request body",
		})
	}

	err = database.DB.Model(&order).Where("id = ?", orderID).Updates(&order).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error updating order",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Order updated successfully",
		"data":    order,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order

	err := database.DB.First(&order, orderID).Delete(&order).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error deleting order",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Order deleted successfully",
	})
}