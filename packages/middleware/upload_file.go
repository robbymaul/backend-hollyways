package middleware

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

// handler function if user input data file
func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			if err == http.ErrMissingFile {
				c.Set("file", "")
				c.Next()
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to upload file",
			})
			return
		}

		// Check if the uploaded file has an allowed image extension
		allowedExtensions := map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".bmp":  true,
			".webp": true,
		}

		// Get the file extension
		fileExtension := filepath.Ext(file.Filename)
		if !allowedExtensions[fileExtension] {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Unsupported file extension",
			})
			c.Set("file", "")
			c.Next()
			return
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Can't open file upload",
			})
			return
		}
		defer src.Close()

		ctx := context.Background()
		var CLOUDE_NAME = os.Getenv("CLOUDE_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		cloudinary, _ := cloudinary.NewFromParams(CLOUDE_NAME, API_KEY, API_SECRET)

		response, err := cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "hollyways"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Set("file", response.SecureURL)
		c.Next()
	}
}
