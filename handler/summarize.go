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

	if err := c.ShouldBindJSON(&req); err != nil || req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request: text is required",
		})
		return
	}

	result, err := service.CallAnthropic(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), // shows real issue (better for debugging)
		})
		return
	}

	c.JSON(http.StatusOK, result)
}