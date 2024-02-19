package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:not null`
	Price       float64 `gorm:not null`
	Description string
}

type User struct {
	gorm.Model
	FirstName string `gorm:not null`
	LastName  string `gorm:not null`
	Email     string `gorm:unique; not null`
	Password  string `gorm:not null`
}

var db *gorm.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHOST + ":" + dbPORT + ")/" + dbNAME + "?parseTime=true"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Product{}, &User{})
}

func GetProducts(c *fiber.Ctx) error {
	var products []Product
	db.Find(&products)
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Product
	result := db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	db.Create(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Product
	result := db.First(&product, productID)
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

	db.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product Product
	result := db.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product Not Found",
		})
	}

	db.Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func main() {

	initDB()
	app := fiber.New()
	app.Get("/products", GetProducts)
	app.Get("/products/:id", GetProduct)
	app.Post("/products", CreateProduct)
	app.Put("/products/:id", UpdateProduct)
	app.Delete("/products/:id", DeleteProduct)
	log.Fatal(app.Listen(":3000"))

}
