package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context)  {
	resData := gin.H{}
	resData["status"] = true
	resData["message"] = "OK"
	c.JSON(http.StatusOK, resData)
	return
}
