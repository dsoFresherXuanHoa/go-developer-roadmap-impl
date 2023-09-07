package services

import (
	"gin-restful-gorm/models"

	"gorm.io/gorm"
)

// Change to Receiver
func FindAllContact(db *gorm.DB, size int, page int) models.Contacts {
	var contacts models.Contacts
	if size == 0 || page == 0 {
		if err := db.Find(&contacts).Error; err != nil {
			return models.Contacts{}
		} else {
			return contacts
		}
	} else {
		// Paging not working right now
		if err := db.Offset(page).Limit(size).Find(&contacts).Error; err != nil {
			return models.Contacts{}
		} else {
			return contacts
		}
	}
}

func FindContactById(id int, db *gorm.DB) models.Contact {
	var contact models.Contact
	if err := db.Where("id = ?", id).First(&contact).Error; err != nil {
		return models.Contact{}
	} else {
		return contact
	}
}

func SaveContact(db *gorm.DB, contact models.Contact) bool {
	if err := db.Create(&contact).Error; err != nil {
		return false
	}
	return true
}

func UpdateContact(db *gorm.DB, id int, contact models.Contact) bool {
	if err := db.Model(&contact).Where("id = ?", id).Updates(models.Contact{Name: contact.Name, Email: contact.Email}).Error; err != nil {
		return false
	}
	return true
}

func DeleteContact(db *gorm.DB, id int) bool {
	var contact models.Contact
	if err := db.Where("id = ?", id).Delete(&contact).Error; err != nil {
		return false
	}
	return true
}
