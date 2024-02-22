package main

import (
	"log"

	Db "githib.com/brewwwwwwww/gofiber-crud/database"
	Handlers "githib.com/brewwwwwwww/gofiber-crud/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	Db.InitDB()
	app := fiber.New()
	app.Get("/products", Handlers.GetProducts)
	app.Get("/products/:id", Handlers.GetProduct)
	app.Post("/products", Handlers.CreateProduct)
	app.Put("/products/:id", Handlers.UpdateProduct)
	app.Delete("/products/:id", Handlers.DeleteProduct)

	app.Post("/users", Handlers.CreateUser)

	log.Fatal(app.Listen(":3000"))

}
