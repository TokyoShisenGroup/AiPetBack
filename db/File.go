package db

import (
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
