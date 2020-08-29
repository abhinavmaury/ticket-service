package routes

import (
	"abhinav/ticket-service/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	router.Use(cors.New(config))

	api := router.Group("/api")
	api.POST("/customer/ticket/generate", handler.GenerateTicket)
	//api.GET("/reset/request", handler.Reset)
	api.POST("/agent/login", handler.LoginController)
	router.GET("/health", handler.HealthCheck)

}
