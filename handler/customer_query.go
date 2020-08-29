package handler

import (
	"abhinav/ticket-service/models"
	"abhinav/ticket-service/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateTicket(c *gin.Context) {
	ctx := context.Background()
	var userTicket models.UserTicket
	if err := c.BindJSON(&userTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         false,
			"message":        err.Error(),
			"additionalInfo": "Unintended JSON data.",
		})
		return
	}
	userTicket.Status = `created`
	if err := userTicket.SaveTicket(&ctx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	agents, _ := models.GetCurrentActiveAgents(context.Background())
	jobQueue := utils.GetJobQueue(agents)
	job := utils.Job{Name: "RESOLVE_QUERY_" + fmt.Sprint(userTicket.TicketID), Fun: models.ProcessTicket, Data: userTicket.TicketID}
	jobQueue.Insert(job)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "task assigned and will be resolved in a while",
	})
}
