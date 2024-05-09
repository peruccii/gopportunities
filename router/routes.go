package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) { // Create Routes
	v1 := router.Group("/api/v1") // Agroup Routes and define api prefix
	{
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.CreateOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.CreateOpeningHandler)
		v1.GET("/openings", handler.CreateOpeningHandler)
	}
}