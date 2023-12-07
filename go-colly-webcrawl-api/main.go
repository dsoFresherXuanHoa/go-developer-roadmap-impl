package main

import (
	"fmt"
	"go-webcrawl-api/src/api/models"
	"go-webcrawl-api/src/api/routes"
	"go-webcrawl-api/src/configs"
)

func main() {
	if database, err := configs.GetGormInstance(); err != nil {
		panic("Can't connect to database via GORM: " + err.Error())
	} else {
		models := []interface{}{&models.News{}}
		database.AutoMigrate(models...)
		fmt.Println("All entity has been synced to database!")
		routes.RouteConfig()
	}
}
