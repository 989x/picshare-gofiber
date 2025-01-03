package controllers

import (
	"os"
	"path/filepath"
	"picshare-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

// Base directories read from .env
var (
	BaseDir         = os.Getenv("UPLOADS_DIR")
	ContentBaseDir  = filepath.Join(BaseDir, "contents")
	BusinessBaseDir = filepath.Join(BaseDir, "businesses")
)

// HandleFileUpload handles multiple files per key and organizes files by public_id
func HandleFileUpload(c *fiber.Ctx, baseDir string) error {
	// Check if BaseDir is set
	if BaseDir == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "UPLOADS_DIR is not set in the environment variables",
		})
	}

	// Ensure directory exists
	if err := ensureDir(baseDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create base directory: " + err.Error(),
		})
	}

	// Generate a public_id
	publicID := utils.GeneratePublicID()

	// Create a directory for the public_id
	publicDir := filepath.Join(baseDir, publicID)
	if err := ensureDir(publicDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create directory for public_id",
		})
	}

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	// Keys to handle
	keys := []string{"cover_image", "body_image"}
	response := fiber.Map{
		"message":   "Images uploaded successfully",
		"public_id": publicID,
	}

	for _, key := range keys {
		// Get files for the key
		files := form.File[key]
		if files == nil {
			continue // Skip if no files for the key
		}

		var filePaths []string
		for _, file := range files {
			// Save each file in the public_id directory
			filePath := filepath.Join(publicDir, file.Filename)
			if err := c.SaveFile(file, filePath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to save image",
				})
			}
			// Record the file path
			filePaths = append(filePaths, "/images/"+filepath.Base(baseDir)+"/"+publicID+"/"+file.Filename)
		}
		// Add file paths to the response
		response[key] = filePaths
	}

	// If no files were uploaded, return an error
	if len(response) == 2 { // Only "message" and "public_id" exist
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No valid files uploaded",
		})
	}

	return c.JSON(response)
}

// ensureDir ensures a directory exists, creating it if necessary
func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
