package main

import (
	"gin-restful-gorm/configs"
	"gin-restful-gorm/routes"
)

func main() {
	if db, err := configs.GormConfig(); err != nil {
		panic("Can't connect to database via GORM!")
	} else {
		routes.RouteConfig(db)
	}
}
