package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ConvId      uint      `gorm:"index;not null;foreignKey:ConvId;references:Conversations(ID)"`
	SenderName  string    `gorm:"index;not null;foreignKey:SenderName;references:User(UserName)"`
	Content     string    `gorm:"not null"`
	CratedAt    time.Time `gorm:"not null"`
	MessageType string    `gorm:"not null"`
	FileId      uint      `gorm:"index;not null;foreignKey:FileId;references:File(ID)"`
}

type MessageGet struct {
	ID          uint
	ConvId      uint
	SenderName  string
	Content     string
	CratedAt    time.Time
	MessageType string
	FileId      uint
}

type MessageRequest struct {
	ID          uint
	ConvId      uint
	SenderName  string
	Content     string
	CratedAt    time.Time
	MessageType string
	FileId      uint
}

type MessageCRUD struct{}

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

func (crud MessageCRUD) GetMessageByConvId(id uint) (*Message, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res Message
	result := db.Where("ConvId = ?", id).First(&res)
	return &res, result.Error
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
	result, err := crud.GetMessageByConvId(id)
	if err != nil {
		return err
	}
	result.Content = "This message has been deleted"
	return crud.UpdateByObject(result)
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
