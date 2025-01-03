package routes

import (
	"picshare-gofiber/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupUploadRoutes sets up routes for handling uploads under api/v1
func SetupUploadRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // Prefix for all routes in this group

	// API: Upload image to 'contents'
	api.Post("/upload/contents", func(c *fiber.Ctx) error {
		return controllers.HandleFileUpload(c, controllers.ContentBaseDir)
	})

	// API: Upload image to 'businesses'
	api.Post("/upload/businesses", func(c *fiber.Ctx) error {
		return controllers.HandleFileUpload(c, controllers.BusinessBaseDir)
	})

	// API: Serve static images
	api.Static("/images", controllers.BaseDir)
}
