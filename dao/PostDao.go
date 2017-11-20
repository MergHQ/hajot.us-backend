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
	if len(post.Content) == 0 {
		return domain.Post{}, errors.New("post not found")
	}
	return domain.Post{Id: post.Id, Content: post.Content, Timestamp: post.CreatedAt}, nil
}

func (dao PostDao) FindNAmount(offset int, amount int) ([]domain.Post, error) {
	var ormPostList []PostORMModel;
	dao.Db.Order("created_at desc").Offset(offset).Limit(amount).Find(&ormPostList)
	if len(ormPostList) == 0 {
		return nil, errors.New("no entries")
	} 
	
	postList := make([]domain.Post, len(ormPostList))
	for i := 0; i < len(ormPostList); i++ {
		postList[i] = domain.Post{Id: ormPostList[i].Id, Content: ormPostList[i].Content, Timestamp: ormPostList[i].CreatedAt}
	}
	return postList, nil
}

func (dao PostDao) Create(content string) domain.Post {
	post := PostORMModel{Content: content}
	dao.Db.Create(&post)
	return domain.Post{Id: post.Id, Content: post.Content, Timestamp: post.CreatedAt}
}