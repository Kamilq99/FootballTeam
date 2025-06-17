package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/buy/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from buy player service",
		})
	})

	router.Run(":8081")
}
