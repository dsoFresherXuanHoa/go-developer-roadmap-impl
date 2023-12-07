package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormInstance *gorm.DB

func GetGormInstance() (*gorm.DB, error) {
	if gormInstance == nil {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Can't load .env files! Check your .env file and try again later: " + err.Error())
			return nil, err
		} else {
			var dns = os.Getenv("GORM_URL")
			if database, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
				fmt.Println("Can't connect to database using GORM: " + err.Error())
				return nil, err
			} else {
				gormInstance = database
				fmt.Println("Connection to database has been created!!!")
			}
		}
	}
	return gormInstance, nil
}
