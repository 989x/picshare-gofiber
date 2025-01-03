package routes

import (
	"os"
	"path/filepath"
	"picshare-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

// SetupUploadRoutes sets up routes for handling uploads
func SetupUploadRoutes(app *fiber.App, baseDir string) {
	contentDir := filepath.Join(baseDir, "contents")
	businessDir := filepath.Join(baseDir, "businesses")

	// Ensure base directories exist
	dirs := []string{contentDir, businessDir}
	for _, dir := range dirs {
		if err := ensureDir(dir); err != nil {
			panic(err)
		}
	}

	// API: Upload image to 'contents'
	app.Post("/upload/contents", func(c *fiber.Ctx) error {
		return handleFileUpload(c, contentDir)
	})

	// API: Upload image to 'businesses'
	app.Post("/upload/businesses", func(c *fiber.Ctx) error {
		return handleFileUpload(c, businessDir)
	})
}

// handleFileUpload handles file upload and organizes files by public_id
func handleFileUpload(c *fiber.Ctx, baseDir string) error {
	// Get the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get image",
		})
	}

	// Generate a public_id using the utils package
	publicID := utils.GeneratePublicID()

	// Create a directory for the public_id
	publicDir := filepath.Join(baseDir, publicID)
	if err := ensureDir(publicDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create directory for public_id",
		})
	}

	// Save the file in the public_id directory
	filePath := filepath.Join(publicDir, file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save image",
		})
	}

	// Respond with the image path
	return c.JSON(fiber.Map{
		"message":   "Image uploaded successfully",
		"public_id": publicID,
		"path":      "/images/" + filepath.Base(baseDir) + "/" + publicID + "/" + file.Filename,
	})
}

// ensureDir ensures a directory exists, creating it if necessary
func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
