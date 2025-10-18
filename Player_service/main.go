package main

import (
	"player_service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/players", controllers.GetPlayers)

	router.Run(":8080")
}
