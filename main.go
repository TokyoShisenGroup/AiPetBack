package main

import (
	"AiPetBack/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化路由
	router.InitRoutes(r)

	r.Run(":8080") // 启动服务器
}
