package db

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	PetName   string    `gorm:"not null"`
	Kind      string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	Age       uint      `gorm:"not null"`
	Birthday  time.Time `gorm:"not null"`
	Weight    float32   `gorm:"not null"`
	OwnerName string    `gorm:"index;not null"`
	User      User      `gorm:"foreignKey:OwnerName;references:UserName"`
}

type GetPet struct{
	PetName   string
	Kind	  string
	Type      string
	Age		  uint
	Birthday  time.Time
	Weight    float32
	OwnerName string
}

type RequestPet struct{
	
}

type PetCRUD struct{}

