package handlers

import (
	"cli-arithmetic-app/app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParseRequest struct {
	Format string `json:"format" binding:"required"`
	Data   []byte `json:"data" binding:"required"`
}
type ComposeRequest struct {
	Format string   `json:"format" binding:"required"`
	Lines  []string `json:"lines" binding:"required"`
}

func ParseHandler(c *gin.Context) {
	var req ParseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lines, err := core.ExecuteParse(req.Data, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"lines": lines})
	}
}

// func ParseFileHandler(c *gin.Context) {

// }

func ComposeHandler(c *gin.Context) {
	var req ComposeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := core.ExecuteCompose(req.Lines, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", bytes)
	}
}

// func ComposeFileHandler(c *gin.Context) {

// }

func ParseHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
