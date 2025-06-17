package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from player service",
		})
	})

	router.Run(":8080")
}
