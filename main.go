package main

import (
	"AiPetBack/chat"
	"AiPetBack/db"
	"AiPetBack/router"
	"AiPetBack/chat/config"
	"AiPetBack/chat/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()

	// 初始化数据库
	initDatabase()

	if config.GetConfig().MsgChannelType.ChannelType == "kafka" {
		kafka.InitProducer(config.GetConfig().MsgChannelType.KafkaTopic, config.GetConfig().MsgChannelType.KafkaHosts)
		kafka.InitConsumer(config.GetConfig().MsgChannelType.KafkaHosts)
		go kafka.ConsumerMsg(chat.ConsumerKafkaMsg)
	}

	// 初始化路由
	router.InitRoutes(ginRouter)

	go chat.MyServer.Start()

	// 启动服务器
	ginRouter.Run(":8080")
}

func initDatabase() {
	dbInstance, err := db.GetDatabaseInstance()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// 自动迁移数据库结构
	dbInstance.AutoMigrate(&db.User{}, &db.Conversation{}, &db.Pet{}, &db.Post{}, &db.Reply{}, &db.Message{})
}
