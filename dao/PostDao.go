package dao

import (
	"github.com/jinzhu/gorm"
	"time"
	"../domain"
	"errors"
)

type PostORMModel struct {
	Id uint `gorm:"primary_key"`
	Content string
	CreatedAt time.Time
}

func (pom PostORMModel) TableName() string {
	return "posts"
}

type PostDao struct {
	Db *gorm.DB
}

func (dao PostDao) Init() {
	dao.Db.AutoMigrate(&PostORMModel{})
}

func (dao PostDao) FindOne(id uint) (domain.Post, error) {
	var post PostORMModel
	dao.Db.First(&post, id)
	if &post == nil {
		return domain.Post{}, errors.New("user not found")
	}
	return domain.Post{Id: post.Id, Content: post.Content, Timestamp: post.CreatedAt}, nil
}
