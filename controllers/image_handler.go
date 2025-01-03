package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"picshare-gofiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	BaseDir         string
	ContentBaseDir  string
	BusinessBaseDir string
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Set base directory from environment
	BaseDir = os.Getenv("BASE_UPLOAD_DIR")
	if BaseDir == "" {
		fmt.Println("BASE_UPLOAD_DIR not set in .env, falling back to default.")
		BaseDir = "/var/www/uploads"
	}

	// Set subdirectories
	ContentBaseDir = filepath.Join(BaseDir, "contents")
	BusinessBaseDir = filepath.Join(BaseDir, "businesses")
}

// HandleFileUpload handles multiple files per key and organizes files by public_id
func HandleFileUpload(c *fiber.Ctx, baseDir string) error {
	// Ensure directory exists
	if err := ensureDir(baseDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create base directory",
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
