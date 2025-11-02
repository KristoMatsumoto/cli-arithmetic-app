package routes

import (
	"cli-arithmetic-app/rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTransformRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/transform/encode", handlers.TransformEncodeHandler)
		api.POST("/transform/encode/chain", handlers.TransformEncodeManyHandler)
		api.POST("/transform/decode", handlers.TransformDecodeHandler)
		api.POST("/transform/decode/chain", handlers.TransformDecodeManyHandler)
		api.GET("/transform/health", handlers.TransformHealthHandler)
	}
}
