package router

import (
	"AiPetBack/chat"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	//log.Logger.Info("newUser", zap.String("newUser", user))
	fmt.Println("newUser", zap.String("newUser", user))
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &chat.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	chat.MyServer.Register <- client
	go client.Read()
	go client.Write()
}

func InitWebsocketRoutes(r *gin.Engine) {
	r.GET("/ws.io", RunSocket)
}