package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Funkcja pomocnicza do tworzenia routera
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/sell/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from sell player service",
		})
	})
	return router
}

func TestSellPlayerEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sell/player", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expected := `{"message":"Hello from sell player service"}`
	assert.JSONEq(t, expected, w.Body.String())
}
