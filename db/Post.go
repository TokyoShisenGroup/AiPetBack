package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"not null"`
	AuthorName  string `gorm:"not null"`
	Content     string `gorm:"type:text;not null"`
	Floor       uint   `gorm:"not null; default:1"`
	IsLocked    bool   `gorm:"default:false"`
	IsDeleted   bool   `gorm:"default:false"`
	IsTop       bool   `gorm:"default:false"`
	IsInvisible bool   `gorm:"default:False"`
}

type PostGet struct {
	ID          uint
	CreatedTime time.Time
	UpdatedTime time.Time
	Title       string
	AuthorName  string
	Avatar      string
	Content     string
	Floor       uint
	IsLocked    bool
}

type PostRequest struct {
	Title    string
	AuthorId uint
	Content  string
}

type PostCRUD struct{}

func (crud PostCRUD) CreateByObject(p *Post) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}

	if p == nil {
		return errors.New("post is nil")
	}
	result := db.Create(p)
	if result.Error != nil {
		return result.Error
	}

	r := &Reply{
		PostId:     p.ID,
		AuthorName: p.AuthorName,
		Content:    p.Content,
		Floor:      1,
	}
	result = db.Create(r)
	return result.Error
}

func (crud PostCRUD) FindById(id uint) (*Post, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var res Post
	result := db.First(&res, id)
	return &res, result.Error
}

func (crud UserCRUD) FindByFuzzyName(name string) ([]Post, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}
	var res []Post
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, result.Error
}

func (crud PostCRUD) FindAll() ([]Post, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var res []Post
	result := db.Find(&res)
	return res, result.Error
}

func (crud PostCRUD) FindAllOrdered() ([]Post, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var top []Post
	var res []Post
	topPosts := db.Where("is_top = ? AND is_invisible = ? AND is_deleted = ?", true, false, false).Find(&top)
	if topPosts.Error != nil {
		return nil, topPosts.Error
	}

	result := db.Order("updated_at desc").Find(&res)
	res = append(res, top...)
	return res, result.Error
}

func (crud PostCRUD) UpdateByObject(p *Post) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}

	result := db.Save(&p)
	return result.Error
}

func (crud PostCRUD) UnsafeDeleteById(id uint) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}

	result := db.Delete(&Post{}, id)
	return result.Error
}

func (crud PostCRUD) SafeDeleteById(id uint) error {
	result, err := crud.FindById(id)
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
