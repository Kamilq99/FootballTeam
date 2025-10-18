package controllers

import (
	"net/http"
	"player_service/models"

	"github.com/gin-gonic/gin"
)

func GetPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Players)
}
