package main

import (
	"gin-restful-gorm/configs"
	"gin-restful-gorm/routes"
)

func main() {
	db := configs.GormConfig()

	routes.RouteConfig(db)
}
