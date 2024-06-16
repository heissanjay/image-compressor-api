package handlers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	c.JSON(200, &gin.H{
		"message": "pong",
	})
}

func HandleCompress(c *gin.Context) {

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing file"})
		return
	}

	file, _ := uploadedFile.Open()
	defer file.Close()
	img, format, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode image"})
		return
	}

	var buf bytes.Buffer
	switch format {
	case "jpeg":
		quality := 75
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(&buf, img)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported image format"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compress image"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=compressed_image."+format)
	c.Header("Content-Length", fmt.Sprint(buf.Len()))

	if _, err := io.Copy(c.Writer, &buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the image"})
		return
	}
}
