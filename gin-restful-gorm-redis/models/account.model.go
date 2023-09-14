package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);NOT NULL" json:"username"`
	Email    string `gorm:"type:varchar(50);" json:"email"`
	Password string `gorm:"type:varchar(50)" json:"password"`
}

type Accounts []Account

type AccountUpdatable struct {
	Email    string `gorm:"type:varchar(50);" json:"email"`
	Password string `gorm:"type:varchar(50)" json:"password"`
}
