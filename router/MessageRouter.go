package router

import (
	"AiPetBack/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var messageCRUD = db.MessageCRUD{}

func RegisterMessagesRoutes(r *gin.Engine){
	r.GET("/messages/:convId", getMessagesByConvId)
	r.POST("/test/messages/create", createMessage)
}

func createMessage(c *gin.Context) {
	var message db.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := messageCRUD.CreateByObject(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, message)
}

func getMessagesByConvId(c *gin.Context) {
	convId := c.Param("convId")
	id, _:=strconv.Atoi(convId)
	messages, err := messageCRUD.GetMessagesByConvId(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

