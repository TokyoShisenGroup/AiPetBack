package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FileName    string    `gorm:"not null"`
	FileType    string    `gorm:"not null"`
	FileSize    int64     `gorm:"not null"`
	FileUrl     string    `gorm:"not null"`
	CreaterName string    `gorm:"index;not null;foreignKey:CreaterName;references:User(UserName)"`
	CreatedTime time.Time `gorm:"not null"`
}

type FileGet struct {
	ID          uint
	FileName    string
	FileType    string
	FileSize    int64
	FileUrl     string
	CreaterName string
	CreatedTime time.Time
}

type FileRequest struct {
	ID          uint
	FileName    string
	FileType    string
	FileSize    int64
	FileUrl     string
	CreaterName string
	CreatedTime time.Time
}

type FileCRUD struct{}

func (crud FileCRUD) CreateByObject(f *File) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if f == nil {
		return errors.New("file not exists")
	}
	result := db.Create(f)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (crud FileCRUD) GetFileByID(id uint) (*File, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res File
	result := db.Where("ID = ?", id).First(&res)
	return &res, result.Error
}

func (crud FileCRUD) UpdateByObject(f *File) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&f)
	return result.Error
}

func (crud FileCRUD) GetFileByCreaterName(name string) ([]File, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []File
	result := db.Where("CreaterName = ?", name).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud FileCRUD) GetFileByFileName(name string) (*File, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res File
	result := db.Where("FileName = ?", name).First(&res)
	return &res, result.Error
}

func (crud FileCRUD) DeleteFileByID(id uint) error {
	result, err := crud.GetFileByID(id)
	if err != nil {
		return err
	}
	result.FileName = "This file has been deleted"
	return crud.UpdateByObject(result)
}

func (crud FileCRUD) GetFileByFuzzyName(name string) ([]File, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []File
	result := db.Where("FileName LIKE ?", "%"+name+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}
