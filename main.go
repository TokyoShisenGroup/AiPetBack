package main

import (
	"AiPetBack/db"
	"AiPetBack/router"
	"AiPetBack/chat"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化数据库
	initDatabase()

	// 初始化路由
	router.InitRoutes(r)
	go chat.MyServer.Start()

	// 启动服务器
	r.Run(":8080")
}

func initDatabase() {
	dbInstance, err := db.GetDatabaseInstance()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// 自动迁移数据库结构
	dbInstance.AutoMigrate(&db.User{}, &db.Conversations{}, &db.Pet{}, &db.Post{}, &db.Reply{})
}
