package db

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `gorm:"not null;unique"`
	Avatar    string ``
	ID        int    `gorm:"primaryKey"`
	PassWord  string `gorm:"not null"`
	IsAdmin   bool   `gorm:"not null;default:false"`
	IsDeleted bool   `gorm:"not null;default:false"`
	LocationX float32
	LocationY float32
}

type UserLogin struct {
	UserName string `json:"UserName" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type UserCRUD struct{}

func (crud UserCRUD) CreateByObject(u *User) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if u == nil {
		return errors.New("User not exists!")
	}
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (crud UserCRUD) GetUserByName(name string) (*User, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var res User
	// 使用 .Where 方法指定列名和查询条件
	result := db.Where("user_name = ?", name).First(&res)
	return &res, result.Error
}

func (crud UserCRUD) UpdateByObject(u *User) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&u)
	return result.Error
}

func (crud UserCRUD) GetUserByFuzzyName(name string) ([]User, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []User
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud UserCRUD) DeleteUserbyName(name string) error {
	result, err := crud.GetUserByName(name)
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

func (crud UserCRUD) GetAllUserOrdered() ([]User, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var top []User
	var res []User
	topUsers := db.Where("is_top = ?", true).Find(&top)
	if topUsers.Error != nil {
		return nil, topUsers.Error
	}

	result := db.Order("updated_at desc").Find(&res)
	res = append(res, top...)
	return res, result.Error
}

func (crud UserCRUD) GetAllUser() ([]User, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var res []User
	result := db.Find(&res)
	return res, result.Error
}
