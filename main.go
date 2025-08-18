package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kaidora-labs/mitter-server/handlers"
)

func main() {
	// TODO: Read Environment Variables

	// TODO: Initialize Database Connection

	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(healthcheck.New())

	// Routes
	userGroup := app.Group("/users")

	userGroup.Get("/", handlers.GetUsersHandler)
	userGroup.Get("/:id", handlers.GetUserHandler)
	userGroup.Post("/", handlers.PostUserHandler)

	// Start the server
	app.Listen(":8080")
}
