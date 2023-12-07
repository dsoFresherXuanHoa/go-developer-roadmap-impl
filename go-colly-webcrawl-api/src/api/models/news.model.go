package models

import "gorm.io/gorm"

type News struct {
	gorm.Model `json:"-"`
	Title      string `json:"-" gorm:"column:title"`
	Content    string `json:"-" gorm:"column:content"`
}

type NewsCreatable struct {
	gorm.Model `json:"-"`
	Title      *string `json:"title" gorm:"column:title"`
	Content    *string `json:"content" gorm:"column:content"`
}

type NewsList []News
type NewsCreatableList []NewsCreatable

func (News) GetTableName() string          { return "news" }
func (NewsCreatable) GetTableName() string { return News{}.GetTableName() }
func (NewsList) GetTableName() string      { return News{}.GetTableName() }
func (NewsList) NewsCreatableList() string { return News{}.GetTableName() }
