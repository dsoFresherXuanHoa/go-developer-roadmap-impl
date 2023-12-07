package entity

import (
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model `json:"-"`
	FullName   string `json:"fullName"`
	Username   string `json:"username"`
	Password   string `json:"-"`
}

type RegisterCreatable struct {
	gorm.Model `json:"-"`
	FullName   *string `json:"fullName"`
	Username   *string `json:"username"`
	Password   *string `json:"password"`
}
