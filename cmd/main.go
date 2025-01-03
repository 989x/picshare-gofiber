package main

import (
	"fmt"
	"log"
	"picshare-gofiber/routes"
	"picshare-gofiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load environment variables
	if err := utils.LoadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Check required environment variables
	utils.MustGetEnv("UPLOADS_DIR") // Ensure UPLOADS_DIR is set

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupUploadRoutes(app)

	// Start the server
	port := "8081"
	fmt.Println("Server running on port " + port)
	app.Listen("0.0.0.0:" + port)
}
