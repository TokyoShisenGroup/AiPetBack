package db

import (
	"errors"
	"AiPetBack/chat/protocol"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ConvId      	uint      `gorm:"index;not null;foreignKey:ConvId;references:Conversations(ID)"`
	SenderName  	string    `gorm:"index;not null;foreignKey:SenderName;references:User(UserId)"`
	ReceiverName 	string    `gorm:"index;not null;foreignKey:ReceiverName;references:User(UserId)"`
	ContentType 	int32     `gorm:"not null"`
	Content     	string    `gorm:"not null"`
	MessageType 	int32     `gorm:"not null"`
	FileUrl      	string     `gorm:"index;not null;foreignKey:FileId;references:File(ID)"`
}

type MessageGet struct {
	gorm.Model
	ConvId      	uint
	SenderName  	string    
	ReceiverName 	string
	ContentType 	int32
	Content     	string
	MessageType 	int32 
	FileUrl      	string
}

type MessageRequest struct {
	gorm.Model
	ConvId      	uint
	SenderName  	string    
	ReceiverName 	string
	ContentType 	int32
	Content     	string
	MessageType 	int32 
	FileUrl      	string
}

type MessageCRUD struct{}

func decodeMessage(m *protocol.Message)*Message{
	msg:=&Message{
		SenderName: m.From,
		ReceiverName: m.To,
		ContentType: m.ContentType,
		Content: m.Content,
		MessageType: m.MessageType,
		FileUrl: m.Url,
	}
	return msg
}

func encodeMessage(m *Message)*protocol.Message{
	msg:=&protocol.Message{
		From: m.SenderName,
		To: m.ReceiverName,
		ContentType: m.ContentType,
		Content: m.Content,
		MessageType: m.MessageType,
		Url: m.FileUrl,
	}
	return msg
}

func SaveMessage(m *protocol.Message)error{
	db,err:=GetDatabaseInstance()
	if err!=nil{
		return err
	}
	msg:=decodeMessage(m)
	result:=db.Create(msg)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

func (crud MessageCRUD) CreateByObject(m *Message) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if m == nil {
		return errors.New("message not exists")
	}
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (crud MessageCRUD) GetMessagesByConvId(id uint) (*protocol.Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res Message
	result := db.Where("ConvId = ?", id).First(&res)

	msg:=encodeMessage(&res)

	return msg, result.Error
}

func (crud MessageCRUD) UpdateByObject(m *Message) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&m)
	return result.Error
}

func (crud MessageCRUD) GetMessageBySenderName(name string) ([]Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Message
	result := db.Where("SenderName = ?", name).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud MessageCRUD) DeleteMessageByID(id uint) error {
	result, err := crud.GetMessagesByConvId(id)
	if err != nil {
		return err
	}
	result.Content = "This message has been deleted"
	return crud.UpdateByObject(decodeMessage(result))
}

func (crud MessageCRUD) GetAllMessageOrdered() ([]Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var res []Message
	result := db.Order("CratedAt desc").Find(&res)
	return res, result.Error
}

func (crud MessageCRUD) GetMessageByFuzzyContent(content string) ([]Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Message
	result := db.Where("content LIKE ?", "%"+content+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud MessageCRUD) GetMessageByFuzzySenderName(name string) ([]Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Message
	result := db.Where("SenderName LIKE ?", "%"+name+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}
