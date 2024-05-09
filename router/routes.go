package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) { // Create Routes
	v1 := router.Group("/api/v1") // Agroup Routes and define api prefix
	{
		v1.GET("/example")
	}
}