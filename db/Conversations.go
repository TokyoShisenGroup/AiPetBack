package db

import (
	"errors"

	"gorm.io/gorm"
)

type Conversations struct {
	gorm.Model
	User1Name string `gorm:"index;not null;foreignKey:User1Name;references:User(UserName)"`
	User2Name string `gorm:"index;not null;foreignKey:User2Name;references:User(UserName)"`
}

type ConversationGet struct {
	ID        uint
	User1Name string
	User2Name string
}

type ConversationRequest struct {
	ID        uint
	User1Name string
	User2Name string
}

type ConversationCRUD struct{}

func (crud ConversationCRUD) CreateByObject(c *Conversations) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("conversation not exists")
	}
	result := db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (crud ConversationCRUD) GetConversationByID(id uint) (*Conversations, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res Conversations
	result := db.Where("ID = ?", id).First(&res)
	return &res, result.Error
}

func (crud ConversationCRUD) UpdateByObject(c *Conversations) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&c)
	return result.Error
}

func (crud ConversationCRUD) GetConversationByUserName(name string) ([]Conversations, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res, res2 []Conversations
	result := db.Where("user1_name = ? OR user2_name = ?", name, name).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}

	res = append(res, res2...)
	return res, result.Error
}

func (crud ConversationCRUD) GetConversationByUsers(user1, user2 string) ([]Conversations, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Conversations
	result := db.Where("(user1_name = ? AND user2_name = ?) OR (user1_name = ? AND user2_name = ?)", user1, user2, user2, user1).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}