package router

import (
	"net/http"
	"strconv"

	"AiPetBack/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var conversationCRUD = db.ConversationCRUD{}

func RegisterConversationRoutes(r *gin.Engine) {
    r.POST("/conversations", createConversation)
    r.GET("/conversations/id/:id", getConversationByID)
    r.GET("/conversations/user/:user", getConversationsBySingleUser)
    r.GET("/conversations/:user1/:user2", getConversationsByUsers)
    r.PUT("/conversations/:id", updateConversation)
}

func createConversation(c *gin.Context) {
    var conversation db.Conversations
    if err := c.ShouldBindJSON(&conversation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := conversationCRUD.CreateByObject(&conversation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, conversation)
}

func getConversationByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
        return
    }

    conversation, err := conversationCRUD.GetConversationByID(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, conversation)
}

func getConversationsBySingleUser(c *gin.Context) {
    user := c.Param("user")
    
    conversations, err := conversationCRUD.GetConversationByUserName(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, conversations)
}

func getConversationsByUsers(c *gin.Context) {
    user1 := c.Param("user1")
    user2 := c.Param("user2")
    
    conversations, err := conversationCRUD.GetConversationByUsers(user1, user2)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, conversations)
}

func updateConversation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
        return
    }

    var conversation db.Conversations
    if err := c.ShouldBindJSON(&conversation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    conversation.ID = uint(id)
    if err := conversationCRUD.UpdateByObject(&conversation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, conversation)
}