package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);NOT NULL" json:"name"`
	Email   string `gorm:"type:varchar(50)" json:"email"`
	Phone   string `gorm:"type:varchar(100);NOT NULL;UNIQUE;" json:"phone"`
	Address string `gorm:"type:text" json:"address"`
}

type Contacts []Contact
