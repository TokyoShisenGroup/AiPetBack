package router

import (
	"net/http"
	"strconv"

	"AiPetBack/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var postCRUD = db.PostCRUD{}

type PostAndRepliesResponse struct {
	Post    db.Post
	Replies []db.Reply
}

func RegisterPostRoutes(r *gin.Engine) {
    r.POST("/posts", createPost)
    r.GET("/posts/:id", getPost)
    r.GET("/posts", getAllPosts)
    r.PUT("/posts/:id", updatePost)
    r.DELETE("/posts/:id", deletePost)
}

func createPost(c *gin.Context) {
    var post db.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := postCRUD.CreateByObject(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, post)
}

func getPost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    post, err := postCRUD.FindById(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

	replies, err := db.ReplyCRUD{}.FindAllByPostId(uint(id))
	if err != nil {
        if err == gorm.ErrRecordNotFound{
            c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        }
        return
    }
	res:=PostAndRepliesResponse{
		Post: *post,
		Replies: replies,
	}

    c.JSON(http.StatusOK, res)
}

func getAllPosts(c *gin.Context) {
    posts, err := postCRUD.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, posts)
}

func updatePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    var post db.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post.ID = uint(id)
    if err := postCRUD.UpdateByObject(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

func deletePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    if err := postCRUD.SafeDeleteById(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}