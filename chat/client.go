package chat

import (
	"AiPetBack/chat/config"
	"AiPetBack/chat/kafka"
	"AiPetBack/chat/constant"
	"AiPetBack/chat/protocol"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		messageType, message, err := c.Conn.ReadMessage()
		if messageType == 1 {
			MyServer.Broadcast <- message
			return
		}
		if err != nil {
			//log.Logger.Error("client read message error", log.Any("client read message error", err.Error()))
			fmt.Println("client read message error", err.Error())
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)
		// pong
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err2 := proto.Marshal(pong)
			if nil != err2 {
				//log.Logger.Error("client marshal message error", log.Any("client marshal message error", err2.Error()))
				fmt.Println("client marshal message error", err2.Error())
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			if config.GetConfig().MsgChannelType.ChannelType == constant.KAFKA {
				kafka.Send(message)
			} else {
				MyServer.Broadcast <- message
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
