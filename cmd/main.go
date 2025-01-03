package main

import (
	"fmt"
	"picshare-gofiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Base directory for uploads
	baseDir := "/var/www/uploads"

	// Setup routes
	routes.SetupUploadRoutes(app, baseDir)

	// API: Serve static images
	app.Static("/images", baseDir)

	// Start the server
	port := "8081"
	fmt.Println("Server running on port " + port)
	app.Listen("0.0.0.0:" + port)
}
