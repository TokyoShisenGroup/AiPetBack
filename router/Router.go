package router

import (
	"AiPetBack/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化所有路由
func InitRoutes(r *gin.Engine) {
	// Initialize all routes
	initPetRoutes(r) // Initialize routes for pets

	//These routers are written by gin and can't work normally
	//initUserRoutes(r)           // Initialize routes for users
	//initConversationRoutes(r)   // Initialize routes for conversations
	//initPostRoutes(r)           // Initialize routes for posts
	//initReplyRoutes(r)          // Initialize routes for replies
	r.GET("/ws", RunSocket) // Initialize routes for websocket
}

// initUserRoutes initializes the user routes.
func initUserRoutes(r *gin.Engine) {
	userCRUD := db.UserCRUD{}

	users := r.Group("/users")
	{
		users.POST("/", func(c *gin.Context) {
			var user db.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := userCRUD.CreateByObject(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		users.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			user, err := userCRUD.GetUserByName(name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		users.PUT("/:name", func(c *gin.Context) {
			name := c.Param("name")
			var user db.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			existingUser, err := userCRUD.GetUserByName(name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			user.ID = existingUser.ID
			if err := userCRUD.UpdateByObject(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		users.DELETE("/:name", func(c *gin.Context) {
			name := c.Param("name")
			if err := userCRUD.DeleteUserbyName(name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		})
	}
}

// initConversationRoutes initializes the conversation routes.
func initConversationRoutes(r *gin.Engine) {
	conversationCRUD := db.ConversationCRUD{}

	conversations := r.Group("/conversations")
	{
		conversations.POST("/", func(c *gin.Context) {
			var conversation db.Conversations
			if err := c.ShouldBindJSON(&conversation); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := conversationCRUD.CreateByObject(&conversation); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, conversation)
		})

		conversations.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
			}
			conversation, err := conversationCRUD.GetConversationByID(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
				return
			}
			c.JSON(http.StatusOK, conversation)
		})

		conversations.PUT("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
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
		})
	}
}

// initPostRoutes initializes the post routes.
func initPostRoutes(r *gin.Engine) {
	postCRUD := db.PostCRUD{}

	posts := r.Group("/posts")
	{
		posts.POST("/", func(c *gin.Context) {
			var post db.Post
			if err := c.ShouldBindJSON(&post); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := postCRUD.CreateByObject(&post); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, post)
		})

		posts.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
			}
			post, err := postCRUD.FindById(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
				return
			}
			c.JSON(http.StatusOK, post)
		})

		posts.PUT("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
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
		})

		posts.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
			}
			if err := postCRUD.SafeDeleteById(uint(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
		})
	}
}

// initReplyRoutes initializes the reply routes.
func initReplyRoutes(r *gin.Engine) {
	replyCRUD := db.ReplyCRUD{}

	replies := r.Group("/replies")
	{
		replies.POST("/", func(c *gin.Context) {
			var reply db.Reply
			if err := c.ShouldBindJSON(&reply); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := replyCRUD.CreateByObject(&reply); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, reply)
		})

		replies.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
			}
			reply, err := replyCRUD.FindById(uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
				return
			}
			c.JSON(http.StatusOK, reply)
		})

		replies.PUT("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
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
		})

		replies.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
			}
			if err := replyCRUD.SafeDeleteById(uint(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Reply deleted"})
		})
	}
}
