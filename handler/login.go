package handler

import (
	"abhinav/ticket-service/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginController(c *gin.Context) {
	ctx := context.Background()

	var agent models.Agent
	if err := c.BindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         false,
			"message":        err.Error(),
			"additionalInfo": "Unintended JSON data.",
		})
		return
	}
	if err := models.IsUsernamePasswordValid(&ctx, agent); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":         false,
			"message":        err.Error(),
		})
		return
	}
	if err := models.SetUserToActiveState(&ctx, agent); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":         false,
			"message":        err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"message":        "logged in successfully",
	})
	return
}
