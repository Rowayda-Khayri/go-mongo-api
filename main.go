// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"go-mongo-api/api/routes"
)

func main() {
	r := gin.Default()

	// Initialize routes
	api.InitializeRoutes(r)

	r.Run(":8080")
}
