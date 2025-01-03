package main

import (
	"log"
	"picshare-gofiber/routes"
	"picshare-gofiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load .env file if exists
	utils.LoadEnv()

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupUploadRoutes(app)

	// Start the server
	port := "8081"
	log.Printf("Starting server on port %s", port)
	app.Listen("0.0.0.0:" + port)
}
