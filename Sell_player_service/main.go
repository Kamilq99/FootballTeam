package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/sell/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from sell player service",
		})
	})

	router.Run(":8082")
}
