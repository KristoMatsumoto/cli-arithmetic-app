package routes

import (
	"cli-arithmetic-app/rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterParseRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/parse", handlers.ParseHandler)
		// api.POST("/parse/file", handlers.ParseFileHandler)
		api.POST("/compose", handlers.ComposeHandler)
		// api.POST("/compose/file", handlers.ComposeFileHandler)
		api.GET("/parse/health", handlers.ParseHealthHandler)
	}
}
