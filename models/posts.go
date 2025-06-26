package models

import (
	"time"
)

type BlogPost struct {
	Id         uint       `json:"id" gorm:"column:id"`
	Title      string     `json:"title" gorm:"column:title"`
	Content    string     `json:"content" gorm:"column:content"`
	Author     string     `json:"author" gorm:"column:author"`
	Created_At *time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type BlogCreateUpdate struct {
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Author  string `json:"author" gorm:"column:author"`
}

func (BlogPost) TableName() string {
	return "blog_post_management"
}
