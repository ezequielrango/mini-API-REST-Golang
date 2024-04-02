package main

import "github.com/gofiber/fiber/v2"
import "github.com/ezequielrango/mini-API-REST-Golang.git/controllers"

func setupRoutes(app *fiber.App) {
	app.Get("/", controllers.Index)

	app.Post("/", controllers.Create)
}