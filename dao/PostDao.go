package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"../domain"
)

type PostORMModel struct {
	gorm.Model
	Id uint `gorm:"primary_key"`
	Content string
	CreatedAt time.Time
}

type PostDao struct {
	db *gorm.DB
}

func (dao PostDao) Init() {
	db, err := gorm.Open("postgres", "")
	dao.db = db
	if err != nil {
		panic("failed to connect to database")
	}
}

func (dao PostDao) FindOne(id uint) domain.Post {
	var post PostORMModel
	dao.db.First(&post, id)
	if post != nil {
		return domain.Post{id: post.Id, content: post.Content, timestamp: post.CreatedAt}
	} else {
		return nil
	}
}
