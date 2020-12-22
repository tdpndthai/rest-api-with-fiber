package handler

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-with-fiber/database"
	"rest-api-with-fiber/model"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Product
	db.Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product
func GetProduct(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	db := database.DB
	var product model.Product
	db.Find(&product, id)
	if product.Title == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
		return
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// CreateProduct new product
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
		return nil
	}
	db.Create(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if product.Title == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
		return nil
	}
	db.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})

}