package routes

import (
	"picshare-gofiber/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupUploadRoutes sets up routes for handling uploads
func SetupUploadRoutes(app *fiber.App) {
	// API: Upload image to 'contents'
	app.Post("/upload/contents", func(c *fiber.Ctx) error {
		return controllers.HandleFileUpload(c, controllers.ContentBaseDir)
	})

	// API: Upload image to 'businesses'
	app.Post("/upload/businesses", func(c *fiber.Ctx) error {
		return controllers.HandleFileUpload(c, controllers.BusinessBaseDir)
	})

	// API: Serve static images
	app.Static("/images", controllers.BaseDir)
}
