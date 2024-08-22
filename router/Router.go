package router

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化所有路由
func InitRoutes(r *gin.Engine) {
	initPetRoutes(r)
}
