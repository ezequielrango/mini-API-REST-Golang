package main

import (
	"github.com/ezequielrango/mini-API-REST-Golang.git/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, flow!")
    })

    app.Listen(":3000")
}