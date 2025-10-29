package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// RegisterTransformRoutes(r)
	RegisterParseRoutes(r)
	RegisterProcessRoutes(r)
}
