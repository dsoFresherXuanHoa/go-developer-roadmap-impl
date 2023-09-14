package services

import (
	"encoding/json"
	"gin-restful-gorm-redis/models"
	"log"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func FindAllAccount(db *gorm.DB) (models.Accounts, error) {
	var accounts models.Accounts
	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	} else {
		return accounts, nil
	}
}

func FindAccountByUsername(username string, db *gorm.DB, redisClient *redis.Client) (*models.Account, error) {
	var account models.Account
	if result, err := redisClient.HGet("accounts", username).Result(); err != nil {
		return nil, err
	} else if result != "" {
		log.Println("REDIS CLIENT: GET - " + account.Username)
		json.Unmarshal([]byte(result), &account)
		return &account, nil
	} else {
		if err := db.Model(models.Account{Username: username}).First(&account).Error; err != nil {
			return nil, err
		} else {
			log.Println("GORM CLIENT: SAVE - " + account.Username)
			return &account, nil
		}
	}
}

func SaveAccount(db *gorm.DB, account models.Account, redisClient *redis.Client) error {
	if err := db.Create(&account).Error; err != nil {
		return err
	} else {
		if data, err := json.Marshal(account); err != nil {
			return err
		} else {
			redisClient.HSet("accounts", account.Username, data)
			log.Println("REDIS CLIENT: SAVE - " + account.Username)
		}
		return nil
	}
}

func UpdateAccount(db *gorm.DB, username string, account models.AccountUpdatable, redisClient *redis.Client) error {
	if err := db.Table("accounts").Model(&account).Where("username = ?", username).Updates(map[string]interface{}{"Email": account.Email, "Password": account.Password}).Error; err != nil {
		return err
	} else {
		if data, err := json.Marshal(account); err != nil {
			return err
		} else {
			redisClient.HSet("accounts", username, data)
			log.Println("REDIS CLIENT: UPDATE - " + username)
		}
		return nil
	}
}

func DeleteAccount(db *gorm.DB, username string, redisClient *redis.Client) error {
	var account models.Account
	if err := db.Where("username = ?", username).Delete(&account).Error; err != nil {
		return err
	} else {
		if result, err := redisClient.HDel("accounts", username).Result(); err != nil || result == 0 {
			return err
		}
		log.Println("REDIS CLIENT: DELETE - " + username)
		return nil
	}
}
