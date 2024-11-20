package api

import (
	"github.com/gin-gonic/gin"
	"simasware.com.br/email-microservice/api/controllers/emailsender"
	"simasware.com.br/email-microservice/api/controllers/healthcheck"
)

func setEndpoints(server *gin.Engine) {
	groups := server.Group("/api")
	groups.GET("healthcheck", healthcheck.HealthCheck)

	emailGroup := groups.Group("emailsender")
	emailGroup.POST("", emailsender.EmailSender)
}

func StartServer() {
	server := gin.Default()
	setEndpoints(server)
	server.Run()
}
