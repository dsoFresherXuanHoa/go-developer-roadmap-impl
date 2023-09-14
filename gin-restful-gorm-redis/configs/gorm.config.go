package configs

import (
	"gin-restful-gorm-redis/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func GormConfig() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		panic("Can't load .env variable!")
	}

	var MYSQL_ROOT_USER = os.Getenv("MYSQL_ROOT_USER")
	var MYSQL_ROOT_PASSWORD = os.Getenv("MYSQL_ROOT_PASSWORD")
	var DATABASE_NAME = os.Getenv("DATABASE_NAME")
	var MYSQL_NETWORK = os.Getenv("MYSQL_NETWORK")

	dns := MYSQL_ROOT_USER + ":" + MYSQL_ROOT_PASSWORD + "@tcp(" + MYSQL_NETWORK + ":3306)/" + DATABASE_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("DNS: " + dns)
	if db, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		db.AutoMigrate(&models.Account{})
		log.Println("Connection to database has been created and all entity has been synced!!!")
		return db, nil
	}
}
