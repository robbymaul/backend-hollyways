package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Cant open file upload",
			})
			return
		}
		defer src.Close()

		buf := make([]byte, 522)
		_, err = src.Read(buf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "file to large max 512kb",
			})
			return
		}

		mimeType := http.DetectContentType(buf)

		allowedMIMETypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/gif":  true,
			"image/bmp":  true,
			"image/webp": true,
		}

		if !allowedMIMETypes[mimeType] {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Unsuported file type",
			})
			return
		}

		tempFile, err := os.CreateTemp("uploads", "image-*.png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Cannot save file directory",
			})
			return
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Cannot copy file to directory",
			})
		}

		data := tempFile.Name()

		fileName := data[8:]

		c.Set("file", fileName)
		c.Next()
	}
}
