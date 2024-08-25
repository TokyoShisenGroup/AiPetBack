package main

import (
	"AiPetBack/chat"
	"AiPetBack/db"
	"AiPetBack/router"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	ginRouter := gin.Default()

	httpRouter := mux.NewRouter()

	// 初始化数据库
	initDatabase()

	// 初始化路由
	router.InitRoutes(ginRouter)
	//router.RegisterConversationRoutes(httpRouter)
	//router.RegisterPostRoutes(httpRouter)
	//router.RegisterUserRoutes(httpRouter)
	//router.RegisterReplyRoutes(httpRouter)
	ginRouter.Any("/", gin.WrapH(httpRouter))

	go chat.MyServer.Start()

	// 启动服务器
	ginRouter.Run(":8081")
}

func initDatabase() {
	dbInstance, err := db.GetDatabaseInstance()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// 自动迁移数据库结构
	dbInstance.AutoMigrate(&db.User{}, &db.Conversations{}, &db.Pet{}, &db.Post{}, &db.Reply{})
}
