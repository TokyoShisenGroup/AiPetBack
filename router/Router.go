package router

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化所有路由
func InitRoutes(r *gin.Engine) {
	// Initialize all routes
	initPetRoutes(r)            // Initialize routes for pets
	RegisterUserRoutes(r)           // Initialize routes for users
	RegisterConversationRoutes(r)   // Initialize routes for conversations
	RegisterPostRoutes(r)           // Initialize routes for posts
	RegisterReplyRoutes(r)          // Initialize routes for replies
	r.GET("/ws", RunSocket)		// Initialize routes for websocket
}