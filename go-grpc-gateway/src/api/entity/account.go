package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model `json:"-"`
	Email      string `json:"email" gorm:"column:email;not null"`
	Password   string `json:"-" gorm:"column:password;not null"`
}

type AccountCreatable struct {
	gorm.Model `json:"-"`
	Email      *string `json:"email" gorm:"column:email;unique;not null"`
	Password   *string `json:"password" gorm:"column:password;not null"`
}

type AccountUpdatable struct {
	gorm.Model `json:"-"`
	Password   *string `json:"password" gorm:"column:password;not null"`
}

type Accounts []Account

func (Account) GetTableName() string          { return "accounts" }
func (AccountCreatable) GetTableName() string { return Account{}.GetTableName() }
func (AccountUpdatable) GetTableName() string { return Account{}.GetTableName() }
func (Accounts) GetTableName() string         { return Account{}.GetTableName() }
