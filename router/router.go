package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Start server
	router.Run(":8080")
}
