package emailsender

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simasware.com.br/email-microservice/api/rabbit"
	"simasware.com.br/email-microservice/models"
)

func EmailSender(context *gin.Context) {
	var request models.SendEmailRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestId = uuid.New().String()
	request.Id = requestId

	rabbit.QueueMail(request)

	context.JSON(http.StatusOK,
		gin.H{
			"sucess":    "Your email was queued",
			"requestId": requestId,
			"dateTime":  time.Now(),
		})

}
