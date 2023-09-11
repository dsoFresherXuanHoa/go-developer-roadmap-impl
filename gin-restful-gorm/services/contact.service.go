package services

import (
	"gin-restful-gorm/models"

	"gorm.io/gorm"
)

func FindAllContact(db *gorm.DB) (models.Contacts, error) {
	var contacts models.Contacts
	if err := db.Find(&contacts).Error; err != nil {
		return nil, err
	} else {
		return contacts, nil
	}
}

func FindContactById(id int, db *gorm.DB) (models.Contact, error) {
	var contact models.Contact
	if err := db.Where("id = ?", id).First(&contact).Error; err != nil {
		return models.Contact{}, err
	} else {
		return contact, nil
	}
}

func SaveContact(db *gorm.DB, contact models.Contact) (interface{}, error) {
	if result := db.Create(&contact); result.Error != nil {
		return nil, result.Error
	} else {
		return result.RowsAffected, nil
	}
}

func UpdateContact(db *gorm.DB, id int, contact models.Contact) (interface{}, error) {
	if result := db.Model(&contact).Where("id = ?", id).Updates(map[string]interface{}{"Name": contact.Name, "Email": contact.Email, "Phone": contact.Phone, "Address": contact.Address}); result.Error != nil {
		return nil, result.Error
	} else {
		return result.RowsAffected, nil
	}
}

func DeleteContact(db *gorm.DB, id int) (interface{}, error) {
	var contact models.Contact
	if result := db.Where("id = ?", id).Delete(&contact); result.Error != nil {
		return nil, result.Error
	} else {
		return result.RowsAffected, nil
	}
}
