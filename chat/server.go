package chat

import (
	"AiPetBack/chat/config"
	"AiPetBack/chat/constant"
	"AiPetBack/chat/protocol"
	"AiPetBack/chat/utils"
	"AiPetBack/db"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

// 消费kafka里面的消息, 然后直接放入go channel中统一进行消费
func ConsumerKafkaMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	//log.Logger.Info("start server", log.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			fmt.Println("new user logged in" + conn.Name)
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			fmt.Println("new user logged out" + conn.Name)
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			// 保存消息只会在存在socket的一个端上进行保存，防止分布式部署后，消息重复问题
			_, exits := s.Clients[msg.From]
			if exits {
				saveMessage(msg)
			}

			if msg.MessageType == constant.MESSAGE_TYPE_USER {
				client, ok := s.Clients[msg.To]
				if ok {
					msgByte, err := proto.Marshal(msg)
					if err == nil {
						client.Send <- msgByte
					}
				}
			}
		}
	}
}

// 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(message *protocol.Message) {
	if message.ContentType == 2 {
		// 普通的文件二进制上传
		fileSuffix := utils.GetFileType(message.File)

		contentType := utils.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := os.WriteFile(config.GetConfig().StaticPath.FilePath+url, message.File, 0666)
		if err != nil {
			//log.Logger.Error("write file error", log.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}

	db.SaveMessage(message)
}
