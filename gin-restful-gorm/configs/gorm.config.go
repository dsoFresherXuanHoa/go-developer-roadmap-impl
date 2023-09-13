package configs

import (
	"gin-restful-gorm/models"
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

	var USER = os.Getenv("USER")
	var PASSWORD = os.Getenv("PASSWORD")
	var DATABASE = os.Getenv("DATABASE")

	dns := USER + ":" + PASSWORD + "@tcp(contactdb:3306)/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		log.Println("Can't connect to database: " + err.Error())
		return nil, err
	} else {
		db.AutoMigrate(&models.Contact{})
		log.Println("Connection to database has been created and all entity has been synced!!!")
		return db, nil
	}
}
