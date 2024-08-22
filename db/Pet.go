package db

import (
	"errors"
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
	OwnerName string    `gorm:"index;not null;foreignKey:OwnerName;references:User(UserName)"`
	IsDeleted bool      `gorm:"not null;default:false"`
}

type GetPet struct {
	PetName   string
	Kind      string
	Type      string
	Age       uint
	Birthday  time.Time
	Weight    float32
	OwnerName string
}

type RequestPet struct {
	PetName   string
	Kind      string
	Type      string
	Age       uint
	Birthday  time.Time
	Weight    float32
	OwnerName string
}

type PetCRUD struct{}

func (crud PetCRUD) CreateByObject(p *Pet) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("Pet not exists")
	}
	result := db.Create(p)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (crud PetCRUD) GetPetByName(name string) (*Pet, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res Pet
	// 使用 .Where 方法指定列名和查询条件
	result := db.Where("user_name = ?", name).First(&res)
	return &res, result.Error
}

func (crud PetCRUD) GetPetByKind(kind string) ([]Pet, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Pet
	result := db.Where("").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud PetCRUD) GetAllPetOrdered() ([]Pet, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var top []Pet
	var res []Pet
	topUsers := db.Where("is_top = ?", true).Find(&top)
	if topUsers.Error != nil {
		return nil, topUsers.Error
	}

	result := db.Order("updated_at desc").Find(&res)
	res = append(res, top...)
	return res, result.Error
}

func (crud PetCRUD) GetPetByOwner(name string) ([]Pet, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Pet
	result := db.Where("Ownername = ?", name).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud PetCRUD) UpdateByObject(p *Pet) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&p)
	return result.Error
}

func (crud PetCRUD) GetPetByFuzzyName(name string) ([]Pet, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Pet
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud PetCRUD) DeletePetbyName(name string) error {
	result, err := crud.GetPetByName(name)
	if err != nil {
		return err
	}
	result.IsDeleted = true
	err = crud.UpdateByObject(result)
	if err != nil {
		return err
	}
	return nil
}
