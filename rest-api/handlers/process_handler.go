package handlers

import (
	"cli-arithmetic-app/app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProcessRequest struct {
	Processor string   `json:"processor" binding:"required"`
	Data      []string `json:"data" binding:"required"`
}

func ProcessHandler(c *gin.Context) {
	var req ProcessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := core.ExecuteProcess(req.Data, req.Processor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"processor": req.Processor, "result": result})
	}
}

func ProcessHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
