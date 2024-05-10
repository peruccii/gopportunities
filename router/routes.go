package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/handler/opening"
	"github.com/peruccii/gopportunities/handler"
	
)

func initializeRoutes(router *gin.Engine) { // Create Routes
	// Initialize handler
	handler.InitHandler()
	v1 := router.Group("/api/v1") // Agroup Routes and define api prefix
	{
		v1.GET("/opening", opening.ShowOpeningHandler)
		v1.DELETE("/opening", opening.DeleteOpeningHandler)
		v1.POST("/opening", opening.CreateOpeningHandler)
		v1.PUT("/opening", opening.UpdateOpeningHandler)
		v1.GET("/openings", opening.ListOpeningHandler)
	}
}