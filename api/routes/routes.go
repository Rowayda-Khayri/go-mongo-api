package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-mongo-api/api/handlers"
)

// InitializeRoutes sets up all routes for the API
func InitializeRoutes(router *gin.Engine) {
	// Basic route for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!!!!!"})
	})

	// Define routes and their corresponding handlers
	router.POST("/signup", handlers.Signup)
}