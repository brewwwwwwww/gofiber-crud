package handlers

import (
	Database "githib.com/brewwwwwwww/gofiber-crud/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetProducts(c *fiber.Ctx) error {
	var products []Database.Product
	Database.Db.Find(&products)
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Database.Product
	result := Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product Database.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	Database.Db.Create(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Database.Product
	result := Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product Not Found",
		})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	Database.Db.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Database.Product
	result := Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product Not Found",
		})
	}

	Database.Db.Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user Database.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.Password = string(hashedPassword)
	Database.Db.Create(&user)
	return c.JSON(user)
}
