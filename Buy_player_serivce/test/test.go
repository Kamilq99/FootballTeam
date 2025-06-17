package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/buy/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from buy player service",
		})
	})
	return router
}

func TestBuyPlayerEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buy/player", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expected := `{"message":"Hello from buy player service"}`
	assert.JSONEq(t, expected, w.Body.String())
}
