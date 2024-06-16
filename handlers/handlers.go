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
	contentType := c.GetHeader("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
		return
	}

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read request body"})
		return
	}

	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode image"})
		return
	}

	var buf bytes.Buffer
	if contentType == "image/jpeg" {
		quality := 75
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	} else if contentType == "image/png" {
		err = png.Encode(&buf, img)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compress image"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=compressed_image."+format)
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", fmt.Sprint(buf.Len()))

	if _, err := io.Copy(c.Writer, &buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the image"})
		return
	}
}
