package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"-" gorm:"column:username"`
	Password   string `json:"-" gorm:"column:password"`
}

type UserCreatable struct {
	gorm.Model `json:"-"`
	Username   *string `json:"username" gorm:"column:username"`
	Password   *string `json:"password" gorm:"column:password"`
}

type UserUpdatable struct {
	gorm.Model `json:"-"`
	Username   *string `json:"username" gorm:"column:username"`
	Password   *string `json:"password" gorm:"column:password"`
}

type Users []User

func (User) GetTableName() string          { return "users" }
func (UserCreatable) GetTableName() string { return "users" }
func (UserUpdatable) GetTableName() string { return "users" }
func (Users) GetTableName() string         { return "users" }
