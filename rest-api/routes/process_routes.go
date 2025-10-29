package routes

import (
	"cli-arithmetic-app/rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProcessRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/process", handlers.ProcessHandler)
		api.GET("/process/health", handlers.ProcessHealthHandler)
	}
}
