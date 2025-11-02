package handlers

import (
	"cli-arithmetic-app/app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransformRequest struct {
	Format string `json:"format" binding:"required"`
	Data   []byte `json:"data" binding:"required"`
}
type TransformManyRequest struct {
	Formats []string `json:"formats" binding:"required"`
	Data    []byte   `json:"data" binding:"required"`
}

func TransformEncodeHandler(c *gin.Context) {
	var req TransformRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := core.ExecuteEncode(req.Data, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"bytes": bytes})
	}
}

func TransformEncodeManyHandler(c *gin.Context) {
	var req TransformManyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := core.ExecuteEncodeMany(req.Data, req.Formats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"bytes": bytes})
	}
}

func TransformDecodeHandler(c *gin.Context) {
	var req TransformRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := core.ExecuteDecode(req.Data, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"bytes": bytes})
	}
}

func TransformDecodeManyHandler(c *gin.Context) {
	var req TransformManyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := core.ExecuteDecodeMany(req.Data, req.Formats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"bytes": bytes})
	}
}

func TransformHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
