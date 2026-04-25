package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Summarizer/service"
)

type Request struct {
	Text string `json:"text"`
}

func Summarize(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	result, err := service.CallAnthropic(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "AI request failed",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}