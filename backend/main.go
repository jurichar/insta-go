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

	// welcome endpoint
	app.Get("/api", welcome)

	// authentication endpoints
	app.Post("/api/register", routes.Register)
	// app.Post("/api/login", routes.Login)

	// user endpoints
	// app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// product endpoints
	app.Post("/api/products", routes.CreateProduct)

	app.Static("/", "../frontend")

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/register.html")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/login.html")
	})

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
