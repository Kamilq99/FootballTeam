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
	router.GET("/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from player service",
		})
	})
	return router
}

func TestPlayerEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/player", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expectedBody := `{"message":"Hello from player service"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
