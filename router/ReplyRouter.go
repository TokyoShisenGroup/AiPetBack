package router

import (
    "net/http"
    "strconv"

    "AiPetBack/db"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

var replyCRUD = db.ReplyCRUD{}

func RegisterReplyRoutes(r *gin.Engine) {
    r.POST("/post/:postid/reply", createReply)
    r.GET("/post/:postid/getreplies", getRepliesOfPostById)
    //r.GET("/post/replies", getAllReplies)
    r.PUT("/post/reply/:id", updateReply)
    r.DELETE("/post/reply/:id", deleteReply)
}

func createReply(c *gin.Context) {
    var reply db.Reply
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id, err := strconv.Atoi(c.Param("postid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Get Post ID failed"})
        return
    }

    reply.PostId = uint(id)
    if err := replyCRUD.CreateByObject(&reply); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, reply)
}

func getRepliesOfPostById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("postid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Get Post ID failed"})
        return
    }

    replies, err := replyCRUD.FindAllByPostId(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, replies)
}

func getAllReplies(c *gin.Context) {
    replies, err := replyCRUD.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, replies)
}

func updateReply(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reply ID"})
        return
    }

    var reply db.Reply
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    reply.ID = uint(id)
    if err := replyCRUD.UpdateByObject(&reply); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reply)
}

func deleteReply(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reply ID"})
        return
    }

    if err := replyCRUD.SafeDeleteById(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}