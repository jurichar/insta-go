package main

import (
	"backend/database"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Post("/api/login", routes.Login)

	// welcome endpoint
	app.Get("/api", welcome)
	// user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// product endpoints
	app.Post("/api/products", routes.CreateProduct)

	app.Static("/", "../frontend")

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
