package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heissanjay/image-compressor-api/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		handlers.HandlePing(c)
	})

	r.POST("/compress", func(c *gin.Context) {
		handlers.HandleCompress(c)
	})

	r.Run()
}
