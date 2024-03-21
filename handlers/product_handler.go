package handlers

import (
	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/models"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	err := database.DB.Order("id").Find(&products).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error fetching products",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Products fetched successfully",
		"data":    products,
	})
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error fetching product",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Product fetched successfully",
		"data":    product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var requestBody models.Product
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing request body",
		})
	}

	errCreate := database.DB.Create(&requestBody).Error

	if errCreate != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error creating product",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Product created successfully",
		"data":    requestBody,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product

	err := c.BodyParser(&product)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Error parsing request body",
		})
	}

	err = database.DB.Model(&product).Where("id = ?", productID).Updates(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error fetching product",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Product updated successfully",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	err := database.DB.First(&product, id).Delete(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error deleting product",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Product deleted successfully",
	})
}