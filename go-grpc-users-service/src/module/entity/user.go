package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	FullName   string `json:"fullName" gorm:"column:full_name;not null"`
}

type UserCreatable struct {
	gorm.Model `json:"-"`
	FullName   string `json:"fullName" gorm:"column:full_name;not null"`
}

type UserUpdatable struct {
	gorm.Model `json:"-"`
	FullName   string `json:"fullName" gorm:"column:full_name;not null"`
}

type Users []User

func (User) GetTableName() string          { return "users" }
func (UserCreatable) GetTableName() string { return User{}.GetTableName() }
func (UserUpdatable) GetTableName() string { return User{}.GetTableName() }
func (Users) GetTableName() string         { return User{}.GetTableName() }
